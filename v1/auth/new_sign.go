package auth

import (
	"net/http"
	"ohsundosun-api/db"
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
	key := c.GetString("userKey")

	var user model.User

	err := db.BaseUser.Get(key, &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &model.DefaultResponse{
			Message: "unauthorized_refresh_token",
		})
		c.Abort()
		return
	}

	isSecure := true
	if os.Getenv("APP_MODE") == "local" {
		isSecure = false
	}

	c.SetCookie("access-token", util.GetAccessToken(&user), 60*30, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
