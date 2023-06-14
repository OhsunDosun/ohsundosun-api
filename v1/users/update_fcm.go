package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// UpdateFCM godoc
// @Tags Users
// @Summary FCM 변경
// @Description FCM 변경
// @Security AppAuth
// @Param request body users.UpdateFCM.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/users/fcm [patch]
func UpdateFCM(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Token string `json:"token" binding:"required" example:"test"`
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

	if err := db.DB.Create(&model.UserToken{
		UserID: user.ID,
		Token:  req.Token,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_insert",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
