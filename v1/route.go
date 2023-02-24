package v1

import (
	"ohsundosun-api/middleware"
	"ohsundosun-api/v1/auth"
	"ohsundosun-api/v1/users"

	"github.com/gin-gonic/gin"
)

func SetRoute(eg *gin.Engine) {
	v1 := eg.Group("/v1", middleware.CheckAppKey())

	auth.SetRoute(v1)
	users.SetRoute(v1)
}
