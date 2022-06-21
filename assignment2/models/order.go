package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ID           uint
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item `json:"items"`
}
