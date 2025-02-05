package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/classify-number", classifyNumbers)

	log.Print("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
