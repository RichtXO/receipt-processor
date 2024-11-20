package points

import (
	"math"
	"receipt-processor/internal/receipt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func TotalPoints(receipt receipt.Receipt) int {
	var total = 0

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			total += 1
		}
	}

	if totalReceipt, err := strconv.ParseFloat(string(receipt.Total), 64); err == nil {
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
		if len(strings.TrimSpace(string(item.ShortDescription)))%3 == 0 {
			if itemPrice, err := strconv.ParseFloat(string(item.Price), 64); err == nil {
				total += int(math.Ceil(itemPrice * 0.2))
			}
		}
	}

	receiptTime, _ := time.Parse("2006-01-02 15:04",
		string(receipt.PurchaseDate)+" "+string(receipt.PurchaseTime))
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
