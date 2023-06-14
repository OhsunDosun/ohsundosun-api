package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// NewPassword godoc
// @Tags Auth
// @Summary 임시 비밀번호 발급
// @Description 임시 비밀번호 발급
// @Security AppAuth
// @Param request body auth.NewPassword.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request, bad_password"
// @Success 404 {object} model.DefaultResponse "not_found_user"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/auth/password/new [post]
func NewPassword(c *gin.Context) {
	type request struct {
		Email string `json:"email" swaggertype:"string" format:"email" binding:"required,email" example:"test@test.com"`
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

	loc, _ := time.LoadLocation("Asia/Seoul")

	password := util.MakeRandomInt(3) + util.MakeRandomString(5)
	now := time.Now()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	userTemporaryPassword := model.UserTemporaryPassword{
		UserID:    user.ID,
		Password:  string(hashPassword),
		CreatedAt: now,
	}

	if err := db.DB.Create(&userTemporaryPassword).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_insert",
		})
		c.Abort()
		return
	}

	util.SendMail(user.Email, "임시 비밀번호 안내드립니다.", "안녕하세요, "+user.Nickname+"님\n요청하신 임시 비밀번호는 다음과 같습니다.\n\n임시 비밀번호: "+password+"\n발급 시간: "+now.In(loc).Format("2006-01-02 15:04:05")+"\n\n임시 비밀번호는 메일이 발송된 시점부터 48시간동안 유효합니다.\n임시 비밀번호로 로그인 하신 후 바로 비밀번호를 변경해주세요.\n\n감사합니다.")

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
