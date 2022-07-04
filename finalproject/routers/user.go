package routers

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {
	user := route.Group("/users")
	user.POST("/register", controllers.CreateUser)
	user.POST("/login", controllers.UserLogin)

	user.Use(middlewares.Authentication())
	user.DELETE("/", controllers.DeleteUser)
	user.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUser)
}
