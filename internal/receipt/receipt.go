// Package receipt receipts provides json datalayer
package receipt

import (
	"github.com/google/uuid"
	"sync"
)

// Receipt represents Receipt Structure
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// Item represents Individual Items
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipts struct {
	data sync.Map
}

func (m *Receipts) Add(id uuid.UUID, receipt Receipt) bool {
	m.data.Store(id, receipt)
	return true
}

func (m *Receipts) Get(id uuid.UUID) (Receipt, bool) {
	if val, ok := m.data.Load(id); ok {
		return val.(Receipt), true
	}
	return Receipt{}, false
}

func (m *Receipts) Update(id uuid.UUID, receipt Receipt) bool {
	if _, ok := m.data.Load(id); ok {
		m.data.Store(id, receipt)
		return true
	}
	return false
}

func (m *Receipts) Remove(id uuid.UUID) bool {
	_, ok := m.data.LoadAndDelete(id)
	return ok
}
