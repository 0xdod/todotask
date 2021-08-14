package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemStore()
	server := NewServer(store)
	log.Printf("Server started locally on port :8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
