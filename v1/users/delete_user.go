package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Tags Users
// @Summary 회원 탈퇴
// @Description 회원 탈퇴
// @Security AppAuth
// @Success 200 {object} model.DefaultResponse "success"
// @Success 401 {object} model.DefaultResponse "unauthorized"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/users/{userId} [delete]
func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	user := c.MustGet("user").(model.User)

	if user.UUID.String() != userId {
		c.JSON(http.StatusUnauthorized, model.DefaultResponse{
			Message: "unauthorized",
		})
		c.Abort()
		return
	}

	active := false
	now := time.Now()

	if err := db.DB.Model(&user).Updates(model.User{
		Active:     &active,
		InActiveAt: &now,
	}).Error; err != nil {
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
