package main

import (
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()

	// Initialise Avail DA client
	daClient := NewAvailDA()
	// sets up a GET API in route /hello that returns the text "World"
	router.POST("/postData", func(c *gin.Context) {
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
