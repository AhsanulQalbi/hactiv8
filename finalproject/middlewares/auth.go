package middlewares

import (
	"finalproject/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": gin.H{
					"error": err.Error(),
					"msg":   "unauthenticated",
				},
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
