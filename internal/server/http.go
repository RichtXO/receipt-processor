package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"receipt-processor/internal/points"
	"receipt-processor/internal/receipt"
)

// httpServer in-memory storage
type httpServer struct {
	Receipts *receipt.Receipts
}

// ProcessReceiptResponse represents process receipt endpoint
type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

// GetPointsResponse represents get points endpoint
type GetPointsResponse struct {
	Points int `json:"points"`
}

func (handler *httpServer) processReceipt(w http.ResponseWriter, r *http.Request) {
	invalid := func() {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"error\": \"Invalid Parameter\"}"))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		invalid()
		return
	}

	var input receipt.Receipt
	err = json.Unmarshal(body, &input)
	if err != nil {
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

	// Adding UUID and receipt
	if !handler.Receipts.Add(id, input) {
		invalid()
		return
	}

	output, err := json.Marshal(response)
	_, _ = w.Write(output)
}

func (handler *httpServer) getPoints(w http.ResponseWriter, r *http.Request) {
	invalid := func() {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("No receipt found for ID provided!"))
	}

	id := mux.Vars(r)["id"]
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		invalid()
		return
	}

	found, exists := handler.Receipts.Get(uuidParsed)
	if !exists {
		invalid()
		return
	}

	response := struct {
		Points int `json:"points"`
	}{
		points.TotalPoints(found),
	}

	output, err := json.Marshal(response)
	if err != nil {
		invalid()
		return
	}
	_, _ = w.Write(output)
}

// NewHTTPServer Initializing Server
func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		Receipts: &receipt.Receipts{},
	}
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", server.processReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", server.getPoints).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
