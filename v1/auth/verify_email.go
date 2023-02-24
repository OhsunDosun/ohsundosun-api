package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// VerifyEmail godoc
// @Tags Auth
// @Summary 이메일 체크 API
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

	query := base.Query{
		{"email": email},
	}

	var result []*model.User

	db.BaseUser.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	if len(result) == 0 {
		c.JSON(http.StatusOK, &model.DataResponse{
			Message: "success",
			Data: &reponse{
				Available: true,
			},
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data: &reponse{
			Available: false,
		},
	})
}
