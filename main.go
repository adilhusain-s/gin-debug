package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define API endpoint
	router.POST("/jobs", func(c *gin.Context) {
		// Parse request body into a struct or map
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process request body and do something with it
		// For example, create a new job
		jobID := createJob(requestBody)

		// Return response with job ID
		c.JSON(http.StatusCreated, gin.H{
			"job_id": jobID,
		})
	})

	// Start server
	router.Run(":8080")
}

func createJob(requestBody map[string]interface{}) string {
	// Code to create a new job based on the request body
	// and return the job ID
	jobID := "123"
	return jobID
}
