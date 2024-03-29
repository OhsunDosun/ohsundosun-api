package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// UpdateMBTI godoc
// @Tags Users
// @Summary MBTI 변경
// @Description MBTI 변경
// @Security AppAuth
// @Param request body users.UpdateMBTI.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/users/mbti [patch]
func UpdateMBTI(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		MBTI string `json:"mbti" enums:"INTJ,INTP,ENTJ,ENTP,INFJ,INFP,ENFJ,ENFP,ISFJ,ISTJ,ESFJ,ESTJ,ISFP,ISTP,ESFP,ESTP" binding:"required" example:"INTP"`
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

	if err := db.DB.Model(&user).Updates(model.User{MBTI: enum.MBTI(req.MBTI)}).Error; err != nil {
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
