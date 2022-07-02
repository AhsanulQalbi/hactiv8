package routers

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(route *gin.Engine) {
	photo := route.Group("/photos")
	photo.Use(middlewares.Authentication())
	photo.POST("/", controllers.CreatePhoto)
	photo.GET("/", controllers.GetAllPhotos)
	photo.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
	photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
}
