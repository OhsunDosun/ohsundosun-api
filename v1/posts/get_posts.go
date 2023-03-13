package posts

import "github.com/gin-gonic/gin"

// GetPosts godoc
// @Tags Posts
// @Summary 게시물 리스트
// @Description 게시물 리스트
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts [get]
func GetPosts(c *gin.Context) {
}
