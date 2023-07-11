package middleware

import (
	"net/http"
	"ohsundosun-api/model"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckAppKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header["App-Key"]

		if !contains(headers, os.Getenv("APP_KEY")) {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_app_key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckKeepAliveKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header["Keep-Alive-Key"]

		if !contains(headers, os.Getenv("KEEP_ALIVE_KEY")) {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_keep_alive_key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
