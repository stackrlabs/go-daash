package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/joho/godotenv"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash"
	"github.com/stackrlabs/go-daash/availda"
)

// Constants
const (
	EigenDaRpcUrl = "disperser-goerli.eigenda.xyz:443"
)

type Job struct {
	Data   []byte
	Layer  daash.DALayer
	ID     string
	Status map[string]any // Human-readable job status
}
type BlobServer struct {
	queue   chan Job
	Daasher *daash.DABuilder
	Jobs    map[string]Job // map of job ID to job
	sync.Mutex
}

func NewBlobServer() *BlobServer {
	return &BlobServer{
		queue:   make(chan Job, 10),
		Jobs:    make(map[string]Job),
		Daasher: daash.NewDABuilder(),
	}
}

func (b *BlobServer) runJobPool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for job := range b.queue {
		go func(job Job) {
			var jobStatus map[string]any
			ids, proofs, err := postToDA(ctx, job.Data, b.Daasher.Clients[job.Layer])
			if err != nil {
				jobStatus = gin.H{
					"status": "failed",
					"error":  err,
				}
			} else {
				if job.Layer == daash.Avail {
					_, extHash := availda.SplitID(ids[0])
					link := fmt.Sprintf("https://goldberg.avail.tools/#/extrinsics/decode/%s", extHash)
					jobStatus = gin.H{
						"status": "Blob daashed and posted to " + string(job.Layer) + " 🏃",
						"ids":    ids,
						"proofs": proofs,
						"link":   link,
					}
				} else {
					jobStatus = gin.H{
						"status": "Blob daashed and posted to " + string(job.Layer) + " 🏃",
						"ids":    ids,
						"proofs": proofs,
					}
				}
			}
			b.Lock()
			b.Jobs[job.ID] = Job{Data: job.Data, Layer: job.Layer, ID: job.ID, Status: jobStatus}
			b.Unlock()

		}(job)
	}
}

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()
	ctx := context.Background()

	envFile, err := godotenv.Read("../../.env") // read from root
	if err != nil {
		fmt.Println("Error reading .env file")

		return
	}
	authToken := envFile["CELESTIA_AUTH_TOKEN"]

	server := NewBlobServer()
	// Initialise all DA clients
	err = server.Daasher.InitClients(ctx, []daash.DALayer{daash.Avail, daash.Celestia, daash.Eigen}, "./avail-config.json", authToken)
	if err != nil {
		fmt.Printf("failed to build DA clients: %v", err)
		return
	}

	router.POST("/:daName", func(c *gin.Context) {
		daName := c.Param("daName")
		daLayer := daash.DALayer(daName)
		if !daash.IsValidDA(daLayer) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("DA %s not found", daName),
			})
			return
		}

		// Get the data in []byte from the request body
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get raw data: %v", err),
			})
			return
		}

		jobID := generateJobID()
		job := Job{Data: data, Layer: daLayer, ID: jobID, Status: gin.H{
			"status": "pending",
			"jobID":  jobID,
		}}
		server.queue <- job
		server.Lock()
		server.Jobs[jobID] = job
		server.Unlock()
		c.JSON(http.StatusOK, job.Status)
	})

	router.GET("/status/:jobID", func(c *gin.Context) {
		jobID := c.Param("jobID")
		job, ok := server.Jobs[jobID]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("job %s not found", jobID),
			})
			return
		}
		c.JSON(http.StatusOK, job.Status)
	})

	go server.runJobPool()
	// The default port is :8080
	router.Run()
}

func postToDA(c context.Context, data []byte, DAClient da.DA) ([]da.ID, []da.Proof, error) {
	daProofs := make([]da.Proof, 1)
	daIDs := make([]da.ID, 1)
	err := backoff.Retry(func() error {
		proofs, ids, err := DAClient.Submit(c, [][]byte{data}, -1)
		if err != nil {
			fmt.Println("post data: ", err)
			return fmt.Errorf("post data: %w", err)
		}
		daProofs = proofs
		daIDs = ids
		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3))
	if err != nil {
		return nil, nil, fmt.Errorf("retry: %w", err)
	}
	return daProofs, daIDs, nil
}

func generateJobID() string {
	b := make([]byte, 16) // Generates a 128-bit (16 bytes) random hex string
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("error generating random hex string: %v", err)
	}
	randomHexString := hex.EncodeToString(b)
	return randomHexString
}
