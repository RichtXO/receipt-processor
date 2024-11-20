package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"receipt-processor/internal/model"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var inMem = make(map[string]int)

func processReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt model.Receipt

	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Generate a random string and insert into hashmap with totalPoints
	id := uuid.New().String()
	inMem[id] = totalPoints(receipt)
	output, _ := json.Marshal(model.ProcessReceiptResponse{Id: id})

	OutputHandler(w, r, output)
}

func totalPoints(receipt model.Receipt) int {
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

func getPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	points, exists := inMem[id]

	if !exists {
		NotFoundHandler(w, r)
		return
	}

	output, _ := json.Marshal(model.GetPointsResponse{Points: points})
	OutputHandler(w, r, output)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
func OutputHandler(w http.ResponseWriter, r *http.Request, output []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", processReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")
	err := http.ListenAndServe(":8010", router)
	if err != nil {
		return
	}
}
