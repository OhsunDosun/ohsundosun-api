package users

import (
	"net/http"
	"strings"
	"time"

	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/deta/deta-go/service/base"
)

// SignUp godoc
// @Tags Users
// @Summary 회원가입 API
// @Description 회원가입
// @Security AppAuth
// @Param request body users.SignUp.request true "query params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 409 {object} model.DefaultResponse "duplicated_email"
// @Success 500 {object} model.DefaultResponse "failed_put"
// @Router /v1/users [post]
func SignUp(c *gin.Context) {
	type request struct {
		Email    string `json:"email" swaggertype:"string" format:"email" binding:"required" example:"test@test.com"`
		Password string `json:"password" binding:"required" example:"1234"`
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

	query := base.Query{
		{"email": req.Email},
	}

	var result []*model.User

	db.BaseUser.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	if len(result) > 0 {
		c.JSON(http.StatusConflict, &model.DefaultResponse{
			Message: "duplicated_email",
		})
		c.Abort()
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	u := &model.User{
		Email:     req.Email,
		Password:  string(hashPassword),
		Nickname:  req.Nickname,
		MBTI:      mbti,
		CreatedAt: time.Now().Unix(),
	}

	_, err = db.BaseUser.Put(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_put",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
