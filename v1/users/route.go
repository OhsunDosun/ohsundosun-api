package users

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/users")
	{
		auth.GET("", middleware.CheckAccessToken())
		auth.POST("", SignUp)
	}
}
