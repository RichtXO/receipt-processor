// Package receipt receipts provides json datalayer
package receipt

import "regexp"

// Receipt represents Receipt Structure
type Receipt struct {
	Retailer     retailer     `json:"retailer"`
	PurchaseDate purchaseDate `json:"purchaseDate"`
	PurchaseTime purchaseTime `json:"purchaseTime"`
	Total        total        `json:"total"`
	Items        itemArr      `json:"items"`
}

type retailer string
type purchaseDate string
type purchaseTime string
type total string
type itemArr []Item

func (v *retailer) verify() (valid bool) {
	m, err := regexp.Match("^[\\w\\s\\-&]+$", []byte(*v))
	if !m || err != nil {
		return
	}
	return true
}

func (v *purchaseDate) verify() (valid bool) {
	m, err := regexp.Match("^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01])$", []byte(*v)) // Date
	if !m || err != nil {
		return
	}
	return true
}

func (v *purchaseTime) verify() (valid bool) {
	m, err := regexp.Match("^([01]\\d|2[0-3]):[0-5]\\d$", []byte(*v))
	if !m || err != nil {
		return
	}
	return true
}

func (v *total) verify() (valid bool) {
	m, err := regexp.Match("^\\d+\\.\\d{2}$", []byte(*v))
	if !m || err != nil {
		return
	}
	return true
}

func (arr *itemArr) verify() (valid bool) {
	for _, v := range *arr {
		if !v.verify() {
			return
		}
	}
	return true
}

func (v *Receipt) Verify() (valid bool) {
	return v.Retailer.verify() &&
		v.PurchaseDate.verify() &&
		v.PurchaseTime.verify() &&
		v.Items.verify() &&
		v.Total.verify()
}

// Item represents Individual Items
type Item struct {
	ShortDescription shortDescription `json:"shortDescription"`
	Price            price            `json:"price"`
}

type shortDescription string
type price string

func (v *shortDescription) verify() (valid bool) {
	m, err := regexp.Match("^[\\w\\s\\-]+$", []byte(*v))
	if !m || err != nil {
		return
	}
	return true
}

func (v *price) verify() (valid bool) {
	m, err := regexp.Match("^\\d+\\.\\d{2}$", []byte(*v))
	if !m || err != nil {
		return
	}
	return true
}

func (v *Item) verify() (valid bool) {
	return v.ShortDescription.verify() && v.Price.verify()
}

// Was thinking previously to store each receipt to make sure no duplicates are being sent
//type Receipts struct {
//	data sync.Map
//}
//
//func (m *Receipts) Add(id uuid.UUID, receipt Receipt) bool {
//	m.data.Store(id, receipt)
//	return true
//}
//
//func (m *Receipts) Get(id uuid.UUID) (Receipt, bool) {
//	if val, ok := m.data.Load(id); ok {
//		return val.(Receipt), true
//	}
//	return Receipt{}, false
//}
//
//func (m *Receipts) Update(id uuid.UUID, receipt Receipt) bool {
//	if _, ok := m.data.Load(id); ok {
//		m.data.Store(id, receipt)
//		return true
//	}
//	return false
//}
//
//func (m *Receipts) Remove(id uuid.UUID) bool {
//	_, ok := m.data.LoadAndDelete(id)
//	return ok
//}
