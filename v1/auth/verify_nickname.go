package auth

import (
	"net/http"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/gin-gonic/gin"
)

// VerifyNickname godoc
// @Tags Auth
// @Summary 닉네임 체크
// @Description 닉네임 체크
// @Security AppAuth
// @Param nickname path string true "Nickname"
// @Success 200 {object} model.DataResponse{data=auth.VerifyNickname.reponse} "success"
// @Router /v1/auth/verify/nickname/{nickname} [get]
func VerifyNickname(c *gin.Context) {
	type reponse struct {
		Available bool `json:"available" binding:"required"`
	}

	nickname := c.Param("nickname")

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data: &reponse{
			Available: util.VerifyNickname(&nickname),
		},
	})
}
