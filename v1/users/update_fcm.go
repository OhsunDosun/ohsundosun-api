package users

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/deta/deta-go/service/base"
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
		FCM string `json:"fcm" binding:"required" example:"test"`
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

	err = deta.BaseUser.Update(user.Key, base.Updates{
		"fcm": util.DeleteDuplicateItem(append(user.FCM, req.FCM)),
	})

	if err != nil {
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
