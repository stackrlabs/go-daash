package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/joho/godotenv"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/stackrlabs/go-daash"
	"github.com/stackrlabs/go-daash/avail"
	availVerify "github.com/stackrlabs/go-daash/avail/verify"
	celestiaVerify "github.com/stackrlabs/go-daash/celestia/verify"
	"github.com/stackrlabs/go-daash/da"
)

type BlobServer struct {
	queue   chan Job
	Daasher *daash.ClientBuilder
	Jobs    map[string]Job // map of job ID to job
	sync.Mutex
}

func NewBlobServer() *BlobServer {
	return &BlobServer{
		queue:   make(chan Job, 10),
		Jobs:    make(map[string]Job),
		Daasher: daash.NewClientBuilder(),
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

	router.GET("/:daName/verify", func(c *gin.Context) {
		daName := c.Param("daName")
		daLayer := daash.DALayer(daName)
		if !daash.IsValidDA(daLayer) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("DA %s not found", daName),
			})
			return
		}

		verifyDA(c, daLayer, server.Daasher)
	})

	go server.runJobPool()
	// The default port is :8080
	router.Run()
}

func postToDA(c context.Context, data []byte, DAClient da.Client) ([]da.ID, []da.Proof, error) {
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

func verifyDA(c *gin.Context, layer daash.DALayer, daasher *daash.ClientBuilder) {
	var success bool
	switch layer {
	case daash.Celestia:
		txHash, ok := c.GetQuery("txHash")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "txHash is required",
			})
			return
		}
		verifier, err := celestiaVerify.NewVerifier(
			chainMetadata["sepolia"]["rpcUrl"],
			celestiaRpcUrl,
			chainMetadata["sepolia"]["blobstreamverifierAddress"],
			chainMetadata["sepolia"]["blobstreamxAddress"],
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to create verifier: %v", err),
			})
			return
		}
		success, err = verifier.VerifyDataAvailable(txHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to verify data: %v", err),
			})
			return
		}
	case daash.Avail:
		blockHeight, ok := c.GetQuery("blockHeight")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "blockHeight is required",
			})
			return
		}
		extIndex, ok := c.GetQuery("extIndex")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "extIndex is required",
			})
			return
		}
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
		verifier, err := availVerify.NewVerifier(
			daasher.Clients[daash.Avail].(*avail.Client),
			chainMetadata["sepolia"]["rpcUrl"],
			chainMetadata["sepolia"]["availBridgeAddress"],
			chainMetadata["sepolia"]["vectorVerifierAddress"],
			chainMetadata["sepolia"]["vectorXAddress"],
			daasher.Clients[daash.Avail].(*avail.Client).Config.Network,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to create verifier: %v", err),
			})
			return
		}
		success, err = verifier.IsDataIncluded(avail.ID{Height: blockHeightUint, ExtIndex: uint32(extIndexUint)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": fmt.Sprintf("error verifying data: %v", err),
			})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("DA %s not supported yet", layer),
		})
		return
	}

	if !success {
		c.JSON(http.StatusOK, gin.H{
			"success": success,
			"message": "data availability cannot be verified onchain!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": "data availability succesfully verified onchain!",
	})

}
