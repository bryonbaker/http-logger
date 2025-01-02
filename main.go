package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define and parse the port flag
	portPtr := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	port := *portPtr

	// Log the startup message
	fmt.Printf("Server is listening on port %d\n", port)

	// Define the HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Capture the current timestamp in RFC3339 format
		timestamp := time.Now().UTC().Format(time.RFC3339)

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v\n", err)
			// Even if there's an error, respond with 200 OK as per requirements
			w.WriteHeader(http.StatusOK)
			return
		}
		defer r.Body.Close()

		// Convert body to string
		bodyStr := string(body)

		// Log the required information to stdout
		fmt.Printf("%s: %s: %s\n", timestamp, r.Method, bodyStr)

		// Respond with 200 OK and empty body
		w.WriteHeader(http.StatusOK)
	})

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		// Log the error and exit if the server fails to start
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

