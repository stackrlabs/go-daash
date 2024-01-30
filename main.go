package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	rpc "github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/share"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()

	// Initialise Avail DA client
	daClient := NewAvailDA()
	ctx := context.Background()
	// Initialise Celestia DA client
	// Read auth token from env
	authToken := os.Getenv("CELESTIA_AUTH_TOKEN")
	if authToken == "" {
		fmt.Println("AUTH_TOKEN is not set")
		return
	}
	client, err := rpc.NewClient(ctx, "http://localhost:26658", authToken)
	if err != nil {
		fmt.Errorf("failed to create rpc client: %v", err)
	}

	// Use random hex for namespace
	nsBytes := make([]byte, 10)
	_, err = hex.Decode(nsBytes, []byte("9cb73e106b03d1050a13"))
	if err != nil {
		log.Fatalln("invalid hex value of a namespace:", err)
	}
	namespace, err := share.NewBlobNamespaceV0(nsBytes)
	celestia := NewCelestiaDA(client, namespace, -1, ctx)

	// Initalise EigenDA client
	eigen := NewEigendaDAClient("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
	fmt.Println(eigen)
	// sets up a GET API in route /hello that returns the text "World"
	router.POST("/Avail", func(c *gin.Context) {
		// Get the data in []byte from the request body
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get raw data: %v", err),
			})
			return
		}
		// Post the data to DA
		err = postToDA(data, daClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("post to DA: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Daaaash",
		})
	})

	router.POST("/Celestia", func(c *gin.Context) {
		// Get the data in []byte from the request body
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get raw data: %v", err),
			})
			return
		}
		// Post the data to DA
		ids, proofs, err := celestia.Submit(c, [][]byte{data}, -1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("post to DA: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Daaaash",
			"ids":     ids,
			"proofs":  proofs,
		})
	})

	router.POST("/Eigen", func(c *gin.Context) {
		// Get the data in []byte from the request body
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get raw data: %v", err),
			})
			return
		}
		// Post the data to DA
		blobInfo, err := eigen.disperseBlob(ctx, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("post to DA: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Daaaash",
			"blobInfo": blobInfo,
		})
	})
	// Run implements a http.ListenAndServe() and takes in an optional Port number
	// The default port is :8080
	router.Run()
}

func postToDA(data []byte, availDAClient DAClient) error {
	err := backoff.Retry(func() error {
		resp, err := availDAClient.PostData(data)
		if err != nil {
			return fmt.Errorf("post data: %w", err)
		}
		fmt.Println(resp)
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return fmt.Errorf("retry: %w", err)
	}
	return nil
}
