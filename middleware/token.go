package middleware

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("access-token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		tokenString := cookie.Value

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		claims := model.TokenClaims{}

		_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		var user model.User

		if err := db.DB.Model(&model.User{}).First(&user, "uuid = UUID_TO_BIN(?)", claims.UUID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func CheckRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("refresh-token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_refresh_token",
			})
			c.Abort()
			return
		}

		tokenString := cookie.Value

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_refresh_token",
			})
			c.Abort()
			return
		}

		claims := model.TokenClaims{}

		_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, model.DefaultResponse{
				Message: "unauthorized_refresh_token",
			})
			c.Abort()
			return
		}

		var user model.User

		if err := db.DB.Model(&model.User{}).First(&user, "uuid = UUID_TO_BIN(?)", claims.UUID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "unauthorized_access_token",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
