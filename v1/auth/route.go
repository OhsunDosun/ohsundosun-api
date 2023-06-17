package auth

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.GET("sign", middleware.CheckAccessToken(), SignCheck)
		auth.POST("sign", SignIn)
		auth.POST("sign/new", middleware.CheckRefreshToken(), NewSign)

		auth.GET("verify/email/:email", VerifyEmail)
		auth.GET("verify/nickname/:nickname", VerifyNickname)
		auth.POST("password/new", NewPassword)
	}
}
