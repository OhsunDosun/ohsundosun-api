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
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/auth/sign/new [post]
func NewSign(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	accessToken := util.GetAccessToken(&user)

	isSecure := os.Getenv("APP_MODE") != "local"

	c.SetCookie("access-token", accessToken, 60*30, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
