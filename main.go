package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "Hello from Golang application to CLO835 class!")
		} else {
			http.Error(w, "Unsupported request method.", http.StatusMethodNotAllowed)
		}
	})

	// Set up the server on localhost:80
	fmt.Println("Starting server on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
