package models

type Order struct {
	ID           uint   `gorm:"primary_key"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item
}
