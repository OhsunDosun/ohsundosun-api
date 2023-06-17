package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignIn godoc
// @Tags Auth
// @Summary 로그인
// @Description 로그인
// @Security AppAuth
// @Param request body auth.SignIn.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request, bad_password"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Router /v1/auth/sign [post]
func SignIn(c *gin.Context) {
	type request struct {
		Type     string  `json:"type" enums:"DEFAULT, NEW_PASSWORD" binding:"required" example:"DEFAULT"`
		Email    string  `json:"email" swaggertype:"string" format:"email" binding:"required,email" example:"test@test.com"`
		Password string  `json:"password" binding:"required,alphanum,min=8,max=16" example:"test1234"`
		Token    *string `json:"token" example:"test"`
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

	var user *model.User

	if err := db.DB.Model(&model.User{}).First(&user, "email = ?", req.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_user",
		})
		c.Abort()
		return
	}

	if strings.ToUpper(req.Type) == "DEFAULT" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

		if err != nil {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}
	} else {
		var userTemporaryPasswords []model.UserTemporaryPassword

		userTemporaryPasswordsSelect := db.DB.Model(&model.UserTemporaryPassword{})
		userTemporaryPasswordsSelect = userTemporaryPasswordsSelect.Where("user_id", user.ID)
		userTemporaryPasswordsSelect = userTemporaryPasswordsSelect.Where("created_at > DATE_ADD(created_at, INTERVAL -2 DAY)")

		userTemporaryPasswordsSelect.Find(&userTemporaryPasswords)

		if len(userTemporaryPasswords) == 0 {
			c.JSON(http.StatusNotFound, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}

		isErrNil := false

		for _, userTemporaryPassword := range userTemporaryPasswords {
			err = bcrypt.CompareHashAndPassword([]byte(userTemporaryPassword.Password), []byte(req.Password))

			if err == nil {
				isErrNil = true
				break
			}
		}

		if !isErrNil {
			c.JSON(http.StatusNotFound, &model.DefaultResponse{
				Message: "bad_password",
			})
			c.Abort()
			return
		}
	}

	if req.Token != nil {
		db.DB.Create(&model.UserToken{
			UserID: user.ID,
			Token:  *req.Token,
		})
	}

	accessToken := util.GetAccessToken(user)
	refreshToken := util.GetRefreshToken(user)

	isSecure := os.Getenv("APP_MODE") != "local"

	c.SetCookie("access-token", accessToken, 60*30, "/", os.Getenv("APP_HOST"), isSecure, true)
	c.SetCookie("refresh-token", refreshToken, 60*60*24*14, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
