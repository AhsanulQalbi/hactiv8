package main

import (
	"assignment2/controllers"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)

	router.Run(":4444")
}
