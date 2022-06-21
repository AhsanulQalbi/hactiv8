package controllers

import (
	"assignment2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	//db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("isi body : %+v\n", order)
	newOrder := models.Order{
		CustomerName: ctx.PostForm("customer_name"),
		OrderedAt:    time.Now().Local().String(),
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"payload": newOrder,
		// "items":   items,
	})

}
