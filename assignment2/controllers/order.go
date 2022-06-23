package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("isi body : %+v\n", order)
	order.OrderedAt = time.Now().String()
	err := db.Create(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})

}

func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	orders := models.Order{}
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
