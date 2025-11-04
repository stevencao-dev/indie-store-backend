package main

import (
	"log"
	"net/http"
	"github.com/stevencao-dev/gobackend/internal/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.GetProducts)

	log.Println("Backend running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
	