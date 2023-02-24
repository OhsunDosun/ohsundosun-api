package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignIn godoc
// @Tags Auth
// @Summary 로그인 API
// @Description 로그인
// @Security AppAuth
// @Param request body auth.SignIn.request true "query params"
// @Success 201 {object} model.DataResponse{data=auth.SignIn.data} "success"
// @Success 400 {object} model.DefaultResponse "bad_request, bad_password"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Router /v1/auth/sign [post]
func SignIn(c *gin.Context) {
	type request struct {
		Email    string `json:"email" swaggertype:"string" format:"email" binding:"required" example:"test@test.com"`
		Password string `json:"password" binding:"required" example:"1234"`
	}

	type data struct {
		AccessToken  string `json:"accessToken" binding:"required"`
		RefreshToken string `json:"refreshToken" binding:"required"`
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

	user := result[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_password",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, &model.DataResponse{
		Message: "success",
		Data: &data{
			AccessToken:  util.GetAccessToken(user),
			RefreshToken: util.GetRefreshToken(user),
		},
	})
}
