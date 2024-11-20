package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"receipt-processor/pkg/model"
	"receipt-processor/pkg/model/receipts"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	store := receipts.NewMemStore()
	receiptsHandler := NewReceiptHandler(&store)

	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", receiptsHandler.processReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", receiptsHandler.getPoints).Methods("GET")
	err := http.ListenAndServe(":8010", router)
	if err != nil {
		return
	}
}

type receiptStore interface {
	Add(id string, receipt receipts.Receipt) bool
	Get(id string) (receipts.Receipt, bool)
	Update(id string, receipt receipts.Receipt) bool
	Remove(id string) bool
}

type ReceiptHandler struct {
	store receiptStore
}

func NewReceiptHandler(s receiptStore) *ReceiptHandler {
	return &ReceiptHandler{store: s}
}

// API calls
func (handler *ReceiptHandler) processReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt receipts.Receipt

	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		InternalServerErrorHandler(w)
		return
	}

	// Generate a random string
	str := receipt.Retailer + receipt.PurchaseDate + receipt.PurchaseTime + receipt.Total
	id := uuid.NewSHA1(uuid.NameSpaceURL, []byte(str)).String()
	output, _ := json.Marshal(model.ProcessReceiptResponse{Id: id})

	// Before adding, check if receipt already exists
	if _, exist := handler.store.Get(id); exist {
		AlreadyReceivedHandler(w, output)
		return
	}
	// Insert uuid and receipt into map
	if !handler.store.Add(id, receipt) {
		InternalServerErrorHandler(w)
		return
	}

	OutputHandler(w, output)
}

func (handler *ReceiptHandler) getPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	receipt, exists := handler.store.Get(id)
	if !exists {
		NotFoundHandler(w)
		return
	}

	output, _ := json.Marshal(model.GetPointsResponse{Points: totalPoints(receipt)})
	OutputHandler(w, output)
}

func totalPoints(receipt receipts.Receipt) int {
	var total = 0

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			total += 1
		}
	}

	if totalReceipt, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		// 50 points if the total is a round dollar amount with no cents.
		if math.Mod(totalReceipt, 1) == 0 {
			total += 50
		}
		// 25 points if the total is a multiple of 0.25
		if math.Mod(totalReceipt, 0.25) == 0 {
			total += 25
		}
	}

	// 5 points for every 2 items on receipt
	total += len(receipt.Items) / 2 * 5

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			if itemPrice, err := strconv.ParseFloat(item.Price, 64); err == nil {
				total += int(math.Ceil(itemPrice * 0.2))
			}
		}
	}

	receiptTime, _ := time.Parse("2006-01-02 15:04", receipt.PurchaseDate+" "+receipt.PurchaseTime)
	// 6 points if day in purchase date is odd
	if receiptTime.Day()%2 == 1 {
		total += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if receiptTime.Hour() >= 14 && receiptTime.Hour() < 16 {
		total += 10
	}

	return total
}

// *** API Response Handlers ***

func InternalServerErrorHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("404 Not Found"))
}
func AlreadyReceivedHandler(w http.ResponseWriter, output []byte) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Receipt already received!\n"))
	_, _ = w.Write(output)
}
func OutputHandler(w http.ResponseWriter, output []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}
