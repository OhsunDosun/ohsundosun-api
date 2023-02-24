package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// NewPassword godoc
// @Tags Auth
// @Summary 비밀번호 재발급 API
// @Description 비밀번호 재발급
// @Security AppAuth
// @Param request body auth.NewPassword.request true "query params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request, bad_password"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Router /v1/auth/password/new [post]
func NewPassword(c *gin.Context) {
	type request struct {
		Email string `json:"email" swaggertype:"string" format:"email" binding:"required" example:"test@test.com"`
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

	query := base.Query{
		{"email": req.Email},
	}

	var result []*model.User

	db.BaseUser.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_user",
		})
		c.Abort()
		return
	}
}
