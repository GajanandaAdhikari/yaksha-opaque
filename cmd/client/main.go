package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// main function to run the client application.
// It sends a GET request to the server's health endpoint and prints the response.
func main() {
    // Define the URL of the server health endpoint to call.
    url := "http://localhost:8080/health"

    // Perform the HTTP GET request.
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error connecting to server: %v", err)
    }
    // Ensure the response body is closed after reading.
    defer resp.Body.Close()

    // Read the response body.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    // Print the response body as string.
    fmt.Printf("Server response: %s\n", body)
}

