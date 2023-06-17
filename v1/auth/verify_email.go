package auth

import (
	"net/http"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/gin-gonic/gin"
)

// VerifyEmail godoc
// @Tags Auth
// @Summary 이메일 체크
// @Description 이메일 체크
// @Security AppAuth
// @Param email path string true "Email"
// @Success 200 {object} model.DataResponse{data=auth.VerifyEmail.reponse} "success"
// @Router /v1/auth/verify/email/{email} [get]
func VerifyEmail(c *gin.Context) {
	type reponse struct {
		Available bool `json:"available" binding:"required"`
	}

	email := c.Param("email")

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data: &reponse{
			Available: util.VerifyEmail(&email),
		},
	})
}
