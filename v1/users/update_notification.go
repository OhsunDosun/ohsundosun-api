package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// UpdateNotification godoc
// @Tags Users
// @Summary 알림 변경
// @Description 알림 변경
// @Security AppAuth
// @Param request body users.UpdateNotification.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/users/notification [patch]
func UpdateNotification(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Notification *bool `json:"notification" binding:"required" example:"true"`
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

	if err := db.DB.Model(&user).Updates(model.User{Notification: req.Notification}).Error; err != nil {
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
