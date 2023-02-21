package v1

import (
	"ohsundosun-api/v1/auth"

	"github.com/gin-gonic/gin"
)

func SetRoute(eg *gin.Engine) {
	v1 := eg.Group("/v1")

	auth.SetRoute(v1)

}
