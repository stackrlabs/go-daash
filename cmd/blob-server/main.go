package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash"
)

// Constants
const (
	EigenDaRpcUrl = "disperser-goerli.eigenda.xyz:443"
)

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

	// Initialise all DA clients
	builder := daash.NewDABuilder()
	err = builder.InitClients(ctx, []daash.DALayer{daash.Avail, daash.Celestia, daash.Eigen}, "./avail-config.json", authToken)
	if err != nil {
		fmt.Printf("failed to build DA clients: %v", err)
		return
	}
	daClients := builder.Clients

	//TODO: Add a job queue with random ID per job

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

		// Post the data to DA
		ids, proofs, err := postToDA(c, data, daClients[daLayer])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("post to DA: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Blob daashed and posted to " + daName + " 🏃",
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
