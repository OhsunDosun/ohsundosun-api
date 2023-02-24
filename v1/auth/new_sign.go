package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/gin-gonic/gin"
)

// NewSign godoc
// @Tags Auth
// @Summary 토큰 재발급 API
// @Description 토큰 재발급
// @Security AppAuth
// @Security RefreshJWTAuth
// @Success 201 {object} model.DataResponse{data=auth.NewSign.data} "success"
// @Router /v1/auth/sign/new [post]
func NewSign(c *gin.Context) {
	type data struct {
		AccessToken string `json:"accessToken" binding:"required"`
	}

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

	c.JSON(http.StatusCreated, &model.DataResponse{
		Message: "success",
		Data: &data{
			AccessToken: util.GetAccessToken(&user),
		},
	})
}
