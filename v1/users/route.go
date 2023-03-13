package users

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/users")
	{
		auth.POST("", SignUp)
		auth.GET("", middleware.CheckAccessToken(), GetUser)
		auth.PATCH("password", middleware.CheckAccessToken(), UpdatePaasword)
		auth.PATCH("nickname", middleware.CheckAccessToken(), UpdateNickname)
		auth.PATCH("mbti", middleware.CheckAccessToken(), UpdateMBTI)
		auth.PATCH("notification", middleware.CheckAccessToken(), UpdateNotification)
		auth.POST("rating", middleware.CheckAccessToken(), AddRating)
	}
}
