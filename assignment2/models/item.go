package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	ID          uint
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
