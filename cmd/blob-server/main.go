package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	rpc "github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/joho/godotenv"
	"github.com/rollkit/go-da"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/celestiada"
	"github.com/stackrlabs/go-daash/eigenda"
)

// Constants
const (
	CelestiaClientUrl = "http://localhost:26658"
	EigenDaRpcUrl     = "disperser-goerli.eigenda.xyz:443"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()
	ctx := context.Background()

	envFile, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("Error reading .env file")

		return
	}

	// Initialise Avail DA client
	avail, err := availda.New()
	if err != nil {
		fmt.Printf("failed to create avail client: %v", err)
	}

	// Initialise Celestia DA client
	// Read auth token from env
	authToken := envFile["CELESTIA_AUTH_TOKEN"]
	if authToken == "" {
		fmt.Println("AUTH_TOKEN is not set")
		return
	}
	client, err := rpc.NewClient(ctx, CelestiaClientUrl, authToken)
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
	celestia := celestiada.NewClient(client, namespace, -1, ctx)

	// Initalise EigenDA client
	eigen, err := eigenda.New(EigenDaRpcUrl, time.Second*90, time.Second*5)
	if err != nil {
		fmt.Printf("failed to create eigen client: %v", err)
	}

	// Map of DA clients
	daClients := map[string]da.DA{
		"avail":    avail,
		"celestia": celestia,
		"eigen":    eigen,
	}

	router.POST("/:daName", func(c *gin.Context) {
		daName := c.Param("daName")

		if _, ok := daClients[daName]; !ok {
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
		ids, proofs, err := postToDA(c, data, daClients[daName])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("post to DA: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "DAaaaSh",
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
