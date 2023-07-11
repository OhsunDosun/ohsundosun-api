package auth

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"os"

	"github.com/gin-gonic/gin"
)

// SignOut godoc
// @Tags Auth
// @Summary 로그아웃
// @Description 로그아웃
// @Security AppAuth
// @Param request body auth.SignOut.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Router /v1/auth/sign [delete]
func SignOut(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Token *string `json:"token" example:"test"`
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

	if req.Token != nil {
		db.DB.Where("user_id", user.ID).Where("token", req.Token).Delete(&model.UserToken{})
	}

	isSecure := os.Getenv("APP_MODE") != "local"

	c.SetCookie("access-token", "", -1, "/", os.Getenv("APP_HOST"), isSecure, true)
	c.SetCookie("refresh-token", "", -1, "/", os.Getenv("APP_HOST"), isSecure, true)

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
