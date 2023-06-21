package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// AddPost godoc
// @Tags Posts
// @Summary 게시글 추가
// @Description 게시글 추가
// @Security AppAuth
// @Param request formData posts.AddPost.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts [post]
func AddPost(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Title   string   `json:"title" binding:"required,max=30" example:"test"`
		Content string   `json:"content" binding:"required,max=6000" example:"test"`
		MBTI    string   `json:"mbti" enums:"INTJ,INTP,ENTJ,ENTP,INFJ,INFP,ENFJ,ENFP,ISFJ,ISTJ,ESFJ,ESTJ,ISFP,ISTP,ESFP,ESTP" binding:"required" example:"INTP"`
		Type    string   `json:"type" enums:"DAILY,LOVE,FRIEND" binding:"required" example:"DAILY"`
		Images  []string `json:"images" binding:"required" exmaple:"[]"`
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

	post := model.Post{
		UserID:  user.ID,
		MBTI:    enum.MBTI(req.MBTI),
		Type:    enum.PostType(req.Type),
		Title:   req.Title,
		Content: req.Content,
		Images:  strings.Join(req.Images, ","),
	}

	if err := db.DB.Create(&post).Error; err != nil {
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
