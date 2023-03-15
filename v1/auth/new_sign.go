package auth

import (
	"net/http"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"os"

	"github.com/gin-gonic/gin"
)

// NewSign godoc
// @Tags Auth
// @Summary 토큰 재발급
// @Description 토큰 재발급
// @Security AppAuth
// @Success 201 {object} model.DataResponse{data=auth.NewSign.data} "success"
// @Router /v1/auth/sign/new [post]
func NewSign(c *gin.Context) {
	type data struct {
		AccessToken string `json:"accessToken" binding:"required"`
	}

	user := c.MustGet("user").(model.User)

	isSecure := true
	if os.Getenv("APP_MODE") == "local" {
		isSecure = false
	}

	accessToken := util.GetAccessToken(&user)

	c.SetCookie("access-token", accessToken, 60*30, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusCreated, &model.DataResponse{
		Message: "success",
		Data: &data{
			AccessToken: accessToken,
		},
	})
}
