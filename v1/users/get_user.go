package users

import (
	"net/http"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Tags Users
// @Summary 회원정보
// @Description 회원정보
// @Security AppAuth
// @Success 200 {object} model.DataResponse{data=users.GetUser.data} "success"
// @Router /v1/users [get]
func GetUser(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type data struct {
		UUID         string `json:"uuid" binding:"required" example:"test"`
		Nickname     string `json:"nickname" binding:"required" example:"test"`
		MBTI         string `json:"mbti" binding:"required" example:"INTP"`
		Notification bool   `json:"notification" binding:"required" example:"true"`
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data: &data{
			UUID:         user.UUID.String(),
			Nickname:     user.Nickname,
			MBTI:         string(user.MBTI),
			Notification: *user.Notification,
		},
	})
}
