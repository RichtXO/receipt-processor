package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"receipt-processor/internal/points"
	"receipt-processor/internal/receipt"
	"sync"
)

// MEMORY in-memory storage
var MEMORY = sync.Map{}

// processReceipt Post Method of
func processReceipt(w http.ResponseWriter, r *http.Request) {
	invalid := func() {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"error\": \"Invalid Parameter\"}"))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		invalid()
		return
	}

	// Extract json from body
	input := receipt.Receipt{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		invalid()
		return
	}

	// Checking if receipt is valid
	if !input.Verify() {
		invalid()
		return
	}

	// Generate UUID
	id := uuid.NewSHA1(uuid.Max, body)
	response := struct {
		ID string `json:"id"`
	}{
		id.String(),
	}

	// Adding UUID and total point
	MEMORY.Store(id, points.TotalPoints(input))

	// Writing output for client
	output, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// getPoints Get Method
func getPoints(w http.ResponseWriter, r *http.Request) {
	invalid := func() {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("No receipt found for ID provided!"))
	}

	// Converting string id to UUID object
	id := mux.Vars(r)["id"]
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		invalid()
		return
	}

	// Seeing if desired receipt is in memory
	found, exists := MEMORY.Load(uuidParsed)
	if !exists {
		invalid()
		return
	}

	response := struct {
		Points int `json:"points"`
	}{
		found.(int),
	}

	// Writing output for client
	output, err := json.Marshal(response)
	if err != nil {
		invalid()
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// NewHTTPServer Initializing Server
func NewHTTPServer(addr string) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
