package models

type Item struct {
	// gorm.Model
	ID          uint `gorm:"primary_key"`
	LineItemId  uint
	ItemCode    string
	Description string
	Quantity    uint
	OrderID     uint
}
