package points

import (
	"receipt-processor/internal/receipt"
	"testing"
)

func getSmallTestReceipt() receipt.Receipt {
	return receipt.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items: []receipt.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"},
		},
	}
}
func getBigTestReceipt() receipt.Receipt {
	return receipt.Receipt{
		Retailer:     "THIS_retailer_1s_AmAzInG",
		PurchaseDate: "2024-11-19",
		PurchaseTime: "15:02",
		Total:        "853.67",
		Items: []receipt.Item{
			{ShortDescription: "Costco Chicken", Price: "5.00"},
			{ShortDescription: "Gatorade", Price: "1.40"},
			{ShortDescription: "ABC", Price: "2.50"},
			{ShortDescription: "   DVCCCC       ", Price: "1.47"},
			{ShortDescription: "A B C", Price: "0.50"},
			{ShortDescription: "QRD", Price: "7.99"},
			{ShortDescription: "Meds", Price: "50.00"},
			{ShortDescription: "Chips", Price: "756.15"},
			{ShortDescription: "Pepsi - 12-oz", Price: "7.89"},
			{ShortDescription: "Dasani", Price: "5.22"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		},
	}
}
func getTargetTestReceipt() receipt.Receipt {
	return receipt.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []receipt.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35:35",
	}
}
func getMarketTestReceipt() receipt.Receipt {
	return receipt.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []receipt.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
}

func TestTotalPointsSmallReceipt(t *testing.T) {
	testReceipt := getSmallTestReceipt()
	want := 15
	result := TotalPoints(testReceipt)
	if want != result {
		t.Errorf("TestTotalPoints() = %d, want %d", result, want)
	}
}

func TestTotalPointsBigReceipt(t *testing.T) {
	testReceipt := getBigTestReceipt()
	want := 80
	result := TotalPoints(testReceipt)
	if want != result {
		t.Errorf("TestTotalPoints() = %d, want %d", result, want)
	}
}

func TestTotalPointsTargetReceipt(t *testing.T) {
	testReceipt := getTargetTestReceipt()
	want := 28
	result := TotalPoints(testReceipt)
	if want != result {
		t.Errorf("TestTotalPoints() = %d, want %d", result, want)
	}
}

func TestTotalPointsMarketReceipt(t *testing.T) {
	testReceipt := getMarketTestReceipt()
	want := 109
	result := TotalPoints(testReceipt)
	if want != result {
		t.Errorf("TestTotalPoints() = %d, want %d", result, want)
	}
}
