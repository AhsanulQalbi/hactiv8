package models

type Item struct {
	// gorm.Model
	ID          uint   `gorm:"primary_key"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
