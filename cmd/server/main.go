package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

// main function to start the server application.
// Sets up the HTTP routes and listens on port 8080.
func main() {
    // Create a new Gin router instance with default middleware (logger and recovery).
    router := gin.Default()

    // Define a simple health check endpoint for clients to verify server status.
    router.GET("/health", func(c *gin.Context) {
        // Respond with JSON indicating the server is running.
        c.JSON(http.StatusOK, gin.H{
            "status": "server is up",
        })
    })

    // Log to console that the server is starting.
    log.Println("Starting server on http://localhost:8080")

    // Start the HTTP server on port 8080. If it fails, log the error and exit.
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

