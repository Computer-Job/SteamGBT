package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r * http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified
	log.Fatal(http.ListenAndServe(port, nil))
}