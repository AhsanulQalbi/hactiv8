package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"msg":     "Failed to Create Order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"msg":     "Order Created Successfully",
	})

}

func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v", orders)
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"result":  orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}
	success := true
	msg := ""
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(ctx.Param("orderId"))
	order.ID = uint(temp)

	for index, _ := range order.Items {
		if order.Items[index].LineItemId != 0 {
			order.Items[index].OrderID = order.ID
			order.Items[index].ID = order.Items[index].LineItemId
		} else {
			success = false
			msg = "lineItemId can't be empty"
		}
	}

	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": success,
			"msg":     msg,
		})
		return
	}
	fmt.Printf("Value Update: %+v\n", order)
	fmt.Println("Order ID : ", ctx.Param("orderId"))
	err := db.Model(&order).Where("id = ?", ctx.Param("orderId")).Updates(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": success,
	})

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}
	err := db.Where("id = ?", ctx.Param("orderId")).Delete(&orders).Error
	if err != nil {
		panic(err)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"msg":     "Order has successfully deleted",
	})
}
