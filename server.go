package main

import (
	"aictions/ai"
	controller_single "aictions/controllers/single"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	// Create default gin router
	router := gin.Default()

	// Create ai connector instance
	aiConnector, err := ai.NewAiConnector("claude")
	if err != nil {
		fmt.Errorf("Failed to create AI connector: %v", err)
		return
	}

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API",
		})
	})

	// Attach single function
	router.POST("/single", func(c *gin.Context) {
		var requestBody struct {
			Query string `json:"query"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if requestBody.Query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query cannot be empty"})
			return
		}

		response, err := controller_single.Serve(struct {
			Ai    *ai.AiConnector
			Query string
		}{
			Ai:    aiConnector,
			Query: requestBody.Query,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	})

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
