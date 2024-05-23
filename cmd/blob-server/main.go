package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash"
	availVerifier "github.com/stackrlabs/go-daash/availda/verify"
	verify "github.com/stackrlabs/go-daash/celestiada/verify"
)

// Constants
const (
	EigenDaRpcUrl             = "disperser-goerli.eigenda.xyz:443"
	ethEndpoint               = "https://sepolia.drpc.org"
	trpcEndpoint              = "https://celestia-mocha-rpc.publicnode.com:443"
	blobstreamverifierAddress = "0x1Bf80E9b8d21ddCCE11b221E1a23781FEb58EB19" // Contract deployed here: https://sepolia.etherscan.io/address/0x1bf80e9b8d21ddcce11b221e1a23781feb58eb19
	blobstreamxAddress        = "0xf0c6429ebab2e7dc6e05dafb61128be21f13cb1e"
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
		go run(ctx, b, job)
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
	_, err = server.Daasher.InitClients(ctx, []daash.DALayer{daash.Avail, daash.Celestia, daash.Eigen}, "./avail-config.json", authToken, "http://localhost:26658")
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

	router.GET("/celestia/verify/:txHash", func(c *gin.Context) {
		daName := c.Param("daName")
		daLayer := daash.DALayer(daName)
		if daLayer != daash.Celestia {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("DA %s not found", daName),
			})
			return
		}
		txHash := c.Param("txHash")
		verifier, err := verify.NewDAVerifier(
			ethEndpoint,
			trpcEndpoint,
			common.HexToAddress(blobstreamverifierAddress),
			common.HexToAddress(blobstreamxAddress),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to create verifier: %v", err),
			})
			return
		}
		success, err := verifier.VerifyDataAvailable(txHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to verify data: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": success,
			"message": "data verified onchain!",
		})
	})

	router.GET("/avail/verify/:blockHeight/:extIndex", func(c *gin.Context) {
		blockHeight := c.Param("blockHeight")
		extIndex := c.Param("extIndex")
		blockHeightUint, err := strconv.ParseUint(blockHeight, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("blockHeight %s is not a valid uint64", blockHeight),
			})
			return
		}
		extIndexUint, err := strconv.ParseUint(extIndex, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("extIndex %s is not a valid uint64", extIndex),
			})
			return
		}
		verifier, err := availVerifier.NewDAVerifier(
			"./avail-config.json",
			ethEndpoint,
			"0x967F7DdC4ec508462231849AE81eeaa68Ad01389", // Avail bridge deployed on Sepolia
			"0x6B26173C8afF316919542df8dA5A57888e398ee1", // Custom Vector verifier contract deployed on Sepolia
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to create verifier: %v", err),
			})
			return
		}
		success, err := verifier.VerifyDataIncluded(blockHeightUint, extIndexUint)
		if !success || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": fmt.Sprintf("failed to verify data: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": success,
			"message": "data verified onchain!",
		})
	})

	go server.runJobPool()
	// The default port is :8080
	router.Run()
}

func postToDA(c context.Context, data []byte, DAClient da.DA) ([]da.ID, []da.Proof, error) {
	daProofs := make([]da.Proof, 1)
	daIDs := make([]da.ID, 1)
	err := backoff.Retry(func() error {
		ids, proofs, err := DAClient.Submit(c, [][]byte{data}, -1)
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
	return daIDs, daProofs, nil
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

func run(ctx context.Context, b *BlobServer, job Job) {
	var jobStatus map[string]any
	ids, proofs, err := postToDA(ctx, job.Data, b.Daasher.Clients[job.Layer])
	if err != nil {
		jobStatus = gin.H{
			"status": "failed",
			"error":  err,
		}
	} else {
		successLink, err := daash.GetExplorerLink(b.Daasher.Clients[job.Layer], ids)
		if err != nil {
			log.Fatalf("cannot get explorer link: %v", err)
		}
		jobStatus = gin.H{
			"status": "Blob daashed and posted to " + string(job.Layer) + " 🏃",
			"ids":    ids,
			"proofs": proofs,
			"link":   successLink,
		}
	}
	b.Lock()
	b.Jobs[job.ID] = Job{Data: job.Data, Layer: job.Layer, ID: job.ID, Status: jobStatus}
	b.Unlock()
}
