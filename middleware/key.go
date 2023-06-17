package middleware

import (
	"net/http"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

func CheckAppKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header["App-Key"]

		if len(headers) < 1 {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_app_key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
