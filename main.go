package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		// timestamp := time.Now().UTC().Format(time.RFC3339)

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

		// Marshal only the body into JSON
		bodyJSON, err := json.Marshal(bodyStr)
		if err != nil {
			log.Printf("Error marshalling body: %v\n", err)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Output the JSON body on a new line
		fmt.Printf("%s\n", bodyJSON)

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
