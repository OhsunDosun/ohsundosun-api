package auth

import (
	"fmt"
	"net/http"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

func SignCheck(c *gin.Context) {
	fmt.Println(c.GetString("userKey"))

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
