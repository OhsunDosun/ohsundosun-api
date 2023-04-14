package auth

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"os"
	"strings"
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignIn godoc
// @Tags Auth
// @Summary 로그인
// @Description 로그인
// @Security AppAuth
// @Param request body auth.SignIn.request true "body params"
// @Success 201 {object} model.DataResponse{data=auth.SignIn.data} "success"
// @Success 400 {object} model.DefaultResponse "bad_request, bad_password"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Router /v1/auth/sign [post]
func SignIn(c *gin.Context) {
	type request struct {
		Type     string  `json:"type" enums:"DEFAULT, NEW_PASSWORD" binding:"required" example:"DEFAULT"`
		Email    string  `json:"email" swaggertype:"string" format:"email" binding:"required,email" example:"test@test.com"`
		Password string  `json:"password" binding:"required,alphanum,min=8,max=16" example:"test1234"`
		FCM      *string `json:"fcm" example:"test"`
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

	deta.BaseUser.Fetch(&base.FetchInput{
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

	if strings.ToUpper(req.Type) == "NEW_PASSWORD" {
		if !user.NewPasswordCreatedAt.Valid || !user.NewPassword.Valid {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}

		now := time.Now()

		if now.Sub(time.Unix(user.NewPasswordCreatedAt.Int64, 0)) > 3*time.Minute {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.NewPassword.String), []byte(req.Password))

		if err != nil {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

		if err != nil {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}
	}

	isSecure := true
	if os.Getenv("APP_MODE") == "local" {
		isSecure = false
	}

	if req.FCM != nil {
		deta.BaseUser.Update(user.Key, base.Updates{
			"fcm": util.DeleteDuplicateItem(append(user.FCM, *req.FCM)),
		})
	}

	accessToken := util.GetAccessToken(user)
	refreshToken := util.GetRefreshToken(user)

	c.SetCookie("access-token", accessToken, 60*30, "/", os.Getenv("APP_HOST"), isSecure, true)
	c.SetCookie("refresh-token", refreshToken, 60*60*24*14, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusCreated, &model.DataResponse{
		Message: "success",
		Data: &data{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
