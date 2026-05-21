package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello from Concourse Golang! 🚀"); err != nil {
		log.Printf("handler write error: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, "ok"); err != nil {
		log.Printf("healthHandler write error: %v", err)
	}
}
