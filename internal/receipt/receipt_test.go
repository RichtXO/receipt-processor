package receipt

import (
	"github.com/google/uuid"
	"testing"
)

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
		Total: "35:35",
	}
}

func getReceipt2() Receipt {
	return Receipt{
		Retailer:     "123",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "asdf", Price: "0.01"},
		},
		Total: "0.01",
	}
}

func TestReceipts_Add(t *testing.T) {
	test := Receipts{}
	ID := uuid.New()

	test.Add(ID, getReceipt())
	_, ok := test.data.Load(ID)
	if !ok {
		t.Error(ID.String() + " not added!")
	}
}

func TestReceipts_Get(t *testing.T) {
	test := Receipts{}
	ID := uuid.New()
	test.Add(ID, getReceipt())

	_, ok := test.Get(ID)
	if !ok {
		t.Error(ID.String() + " not found!")
	}
}

func TestReceipts_Update(t *testing.T) {
	test := Receipts{}
	ID := uuid.New()
	test.Add(ID, getReceipt())

	ok := test.Update(ID, getReceipt())
	if !ok {
		t.Error(ID.String() + " not being able to update!")
	}
}

func TestReceipts_Remove(t *testing.T) {
	test := Receipts{}
	ID := uuid.New()
	test.Add(ID, getReceipt())

	ok := test.Remove(ID)
	if !ok {
		t.Error(ID.String() + " not removed!")
	}
}
