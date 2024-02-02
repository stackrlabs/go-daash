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
	"github.com/rollkit/go-da"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()

	// Initialise Avail DA client
	avail, err := NewAvailDA()
	if err != nil {
		fmt.Printf("failed to create avail client: %v", err)
	}

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
		fmt.Printf("failed to create rpc client: %v", err)
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
	eigen, err := NewEigendaDAClient("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
	if err != nil {
		fmt.Printf("failed to create eigen client: %v", err)
	}
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
		ids, proofs, err := postToDA(c, data, avail)
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
		ids, proofs, err := postToDA(c, data, celestia)
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
		ids, proofs, err := postToDA(c, data, eigen)
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
	// Run implements a http.ListenAndServe() and takes in an optional Port number
	// The default port is :8080
	router.Run()
}

func postToDA(c *gin.Context, data []byte, DAClient da.DA) ([]da.ID, []da.Proof, error) {
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
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, nil, fmt.Errorf("retry: %w", err)
	}
	return daProofs, daIDs, nil
}
