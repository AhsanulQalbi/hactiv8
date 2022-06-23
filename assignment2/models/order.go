package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	// gorm.Model
	ID           uint `gorm:"primary_key"`
	CustomerName string
	OrderedAt    string
	Items        []Item
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Order Before Create")
	if o.OrderedAt == "" {
		o.OrderedAt = time.Now().String()
	}

	errorMsg := ""

	if len(o.CustomerName) == 0 {
		errorMsg = errorMsg + "customerName can't be empty. "
	}

	for i, _ := range o.Items {
		if len(o.Items[i].ItemCode) == 0 {
			errorMsg = errorMsg + "itemCode item #" + strconv.Itoa(i+1) + " can't be empty. "
		}
	}

	if errorMsg != "" {
		err = errors.New(errorMsg)
	}

	return
}

func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Order Before Update")

	errorMsg := ""
	for index, _ := range order.Items {
		if order.Items[index].LineItemId != 0 {
			order.Items[index].OrderID = order.ID
			order.Items[index].ID = order.Items[index].LineItemId
		} else {
			errorMsg = errorMsg + "lineItemId item #" + strconv.Itoa(index+1) + " can't be empty. "
		}
	}

	if errorMsg != "" {
		err = errors.New(errorMsg)
	}

	return
}
