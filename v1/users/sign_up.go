package users

import (
	"net/http"
	"strings"
	"time"

	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"ohsundosun-api/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignUp godoc
// @Tags Users
// @Summary 회원가입
// @Description 회원가입
// @Security AppAuth
// @Param request body users.SignUp.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 409 {object} model.DefaultResponse "duplicated_email, duplicated_nickname"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/users [post]
func SignUp(c *gin.Context) {
	type request struct {
		Email    string `json:"email" swaggertype:"string" format:"email" binding:"required,email" example:"test@test.com"`
		Password string `json:"password" binding:"required,alphanum,min=8,max=16" example:"test1234"`
		Nickname string `json:"nickname" binding:"required" example:"test"`
		MBTI     string `json:"mbti" enums:"INTJ,INTP,ENTJ,ENTP,INFJ,INFP,ENFJ,ENFP,ISFJ,ISTJ,ESFJ,ESTJ,ISFP,ISTP,ESFP,ESTP" binding:"required" example:"INTP"`
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

	mbti := enum.StringToMBTI(strings.ToUpper(req.MBTI))
	if mbti == 0 {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	if !util.VerifyEmail(&req.Email) {
		c.JSON(http.StatusConflict, &model.DefaultResponse{
			Message: "duplicated_email",
		})
		c.Abort()
		return
	}

	if !util.VerifyNickname(&req.Nickname) {
		c.JSON(http.StatusConflict, &model.DefaultResponse{
			Message: "duplicated_nickname",
		})
		c.Abort()
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	u := &model.User{
		Key:          util.NewULID().String(),
		Email:        req.Email,
		Password:     string(hashPassword),
		Nickname:     req.Nickname,
		MBTI:         mbti,
		CreatedAt:    time.Now().Unix(),
		Notification: true,
		Active:       true,
	}

	_, err = db.BaseUser.Insert(u)

	if err != nil {
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
