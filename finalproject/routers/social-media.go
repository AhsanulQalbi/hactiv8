package routers

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(route *gin.Engine) {
	sosmed := route.Group("/socialmedias")
	sosmed.Use(middlewares.Authentication())
	sosmed.GET("/", controllers.GetAllSocialMedia)
	sosmed.POST("/", controllers.CreateSocialMedia)
	sosmed.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
	sosmed.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
}
