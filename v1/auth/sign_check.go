package auth

import (
	"fmt"
	"net/http"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

func SignCheck(c *gin.Context) {
	fmt.Println(c.Get("user"))

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
