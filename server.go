package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	// Create default gin router
	router := gin.Default()

	//

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API",
		})
	})

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
