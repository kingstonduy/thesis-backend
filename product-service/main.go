package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler for the GET request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is GET
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Send a response
		fmt.Fprintf(w, "Hello, you've made a GET request!")
	})

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
