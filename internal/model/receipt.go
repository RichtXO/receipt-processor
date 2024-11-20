package model

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
