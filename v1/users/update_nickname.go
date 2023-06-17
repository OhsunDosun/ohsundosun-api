package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// UpdateNickname godoc
// @Tags Users
// @Summary 닉네임 변경
// @Description 닉네임 변경
// @Security AppAuth
// @Param request body users.UpdateNickname.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/users/nickname [patch]
func UpdateNickname(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Nickname string `json:"nickname" binding:"required" example:"test"`
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

	if err := db.DB.Model(&user).Updates(model.User{Nickname: req.Nickname}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
