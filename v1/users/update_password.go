package users

import (
	"net/http"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// UpdatePaasword godoc
// @Tags Users
// @Summary 비밀번호 변경
// @Description 비밀번호 변경
// @Security AppAuth
// @Param request body users.UpdatePaasword.request true "query params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Router /v1/users/password [patch]
func UpdatePaasword(c *gin.Context) {
	type request struct {
		OldPassword string `json:"oldPassword" binding:"required" example:"1234"`
		NewPassword string `json:"newPassword" binding:"required" example:"1234"`
	}

	req := &request{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}
}
