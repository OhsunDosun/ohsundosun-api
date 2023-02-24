package middleware

import (
	"net/http"
	"ohsundosun-api/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header["Access-Token"]

		if len(headers) < 1 {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		tokenString := headers[0]

		claims := model.TokenClaims{}

		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		c.Set("userKey", claims.Key)

		c.Next()
	}
}

func CheckRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header["Refresh-Token"]

		if len(headers) < 1 {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_refresh_token",
			})
			c.Abort()
			return
		}

		tokenString := headers[0]

		claims := model.TokenClaims{}

		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_refresh_token",
			})
			c.Abort()
			return
		}

		c.Set("userKey", claims.Key)

		c.Next()
	}
}
