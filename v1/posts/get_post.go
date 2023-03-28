package posts

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// GetPost godoc
// @Tags Posts
// @Summary 게시물 상세
// @Description 게시물 상세
// @Security AppAuth
// @Success 200 {object} model.DataResponse{data=posts.GetPost.data} "success"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Router /v1/posts/{postId} [get]
func GetPost(c *gin.Context) {
	postId := c.Param("postId")

	type data struct {
		Key          string   `json:"key"  binding:"required" example:"test"`
		MBTI         string   `json:"mbti" binding:"required" example:"INTP"`
		Type         string   `json:"type"  binding:"required" example:"DAILY"`
		Nickname     string   `json:"nickname"  binding:"required" example:"test"`
		Title        string   `json:"title"  binding:"required" example:"test"`
		Content      string   `json:"content"  binding:"required" example:"test"`
		Images       []string `json:"images"  binding:"required" example:"test.png,test.png"`
		CreatedAt    int64    `json:"createdAt"  binding:"required"`
		LikeCount    int8     `json:"likeCount" binding:"required" example:"0"`
		CommentCount int8     `json:"commentCount" binding:"required" example:"0"`
	}

	var post model.Post

	err := deta.BasePost.Get(postId, &post)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data: &data{
			Key:          post.Key,
			Nickname:     post.Nickname,
			MBTI:         post.MBTI.String(),
			Title:        post.Title,
			Content:      post.Content,
			Type:         post.Type.String(),
			Images:       post.Images,
			CreatedAt:    post.CreatedAt,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
		},
	})
}
