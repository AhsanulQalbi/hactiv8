package main

import (
	"assignment2/controllers"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	// personController := controller.NewPersonController(db)

	router := gin.Default()

	router.POST("/create-order", controllers.CreateOrder)

	router.Run(":4444")
}
