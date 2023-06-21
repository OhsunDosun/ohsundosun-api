package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// BlockUser godoc
// @Tags Users
// @Summary 회원 차단
// @Description 회원 차단
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/users/{userId}/block [post]
func BlockUser(c *gin.Context) {
	userId := c.Param("userId")

	user := c.MustGet("user").(model.User)

	var blockUser model.User

	if err := db.DB.Model(&model.User{}).First(&blockUser, "uuid = UUID_TO_BIN(?)", userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_user",
		})
		c.Abort()
		return
	}

	block := model.UserBlock{
		UserID:  user.ID,
		BlockID: blockUser.ID,
	}

	if err := db.DB.Create(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_insert",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
