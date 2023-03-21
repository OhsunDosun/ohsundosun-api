package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddPost godoc
// @Tags Posts
// @Summary 게시물 추가
// @Description 게시물 추가
// @Security AppAuth
// @Param request body posts.AddPost.request true "query params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 500 {object} model.DefaultResponse "failed_put"
// @Router /v1/posts [post]
func AddPost(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Title   string   `json:"title" binding:"required,max=30" example:"test"`
		Content string   `json:"content" binding:"required,max=6000" example:"test"`
		Type    string   `json:"type" enums:"DAILY,LOVE,FRIEND" binding:"required" example:"DAILY"`
		Images  []string `json:"images" binding:"required" example:"test.png,test.png"`
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

	postType := enum.StringToPostType(strings.ToUpper(req.Type))
	if postType == 0 {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	p := &model.Post{
		Key:       uuid.New().String(),
		UserKey:   user.Key,
		Nickname:  user.Nickname,
		MBTI:      user.MBTI,
		Title:     req.Title,
		Content:   req.Content,
		Type:      postType,
		Images:    req.Images,
		CreatedAt: time.Now().Unix(),
		Active:    true,
	}

	_, err = db.BasePost.Put(p)

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
