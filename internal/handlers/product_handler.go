package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stevencao-dev/gobackend/internal/models"
)

// Dummy Data for testing ppurposes
var products = []models.Product{
	{ID: 1, Name: "Path of Coding", Price: 999.99, Stock: 10},
	{ID: 2, Name: "ElderPaper III", Price: 499.99, Stock: 25},
	{ID: 3, Name: "GTA VIII: This is the one", Price: 299.99, Stock: 15},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// products stored in JSON format, set type to JSON
	w.Header().Set("Content-Type", "application/json")

	// actually encoding to JSON to send to client
	json.NewEncoder(w).Encode(products)
}