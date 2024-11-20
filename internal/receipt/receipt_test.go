package receipt

import "testing"

func getReceipt() Receipt {
	return Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}
}

func getWrongReceipt() Receipt {
	return Receipt{
		Retailer:     "123",
		PurchaseDate: "NOTaDATE",
		PurchaseTime: "NOTaTIME",
		Items: []Item{
			{ShortDescription: "asdf", Price: "NOTprice"},
		},
		Total: "0.01",
	}
}

func TestCorrectReceipt_Verify(t *testing.T) {
	receipt := getReceipt()
	if !receipt.Verify() {
		t.Error("Correct Receipt does not verify!")
	}
}

func TestWrongReceipt_Verify(t *testing.T) {
	receipt := getWrongReceipt()
	if receipt.Verify() {
		t.Error("Wrong Receipt does not verify!")
	}
}
