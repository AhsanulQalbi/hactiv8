package routers

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRouter(route *gin.Engine) {
	comment := route.Group("/comments")
	comment.Use(middlewares.Authentication())
	comment.POST("/", controllers.CreateComment)
	comment.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
	comment.GET("/", controllers.GetAllComments)
	comment.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
}
