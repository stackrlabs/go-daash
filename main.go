package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()

	// sets up a GET API in route /hello that returns the text "World"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Daaaash",
		})
	})

	// Run implements a http.ListenAndServe() and takes in an optional Port number
	// The default port is :8080
	router.Run()
}
