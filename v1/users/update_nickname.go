package users

import "github.com/gin-gonic/gin"

// UpdateNickname godoc
// @Tags Users
// @Summary 닉네임 변경
// @Description 닉네임 변경
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/users/nickname [patch]
func UpdateNickname(c *gin.Context) {
}
