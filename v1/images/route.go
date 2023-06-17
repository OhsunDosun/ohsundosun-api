package images

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/images")
	{
		auth.POST("", middleware.CheckAccessToken(), AddImage)
	}
}
