package posts

import "github.com/gin-gonic/gin"

// GetPost godoc
// @Tags Posts
// @Summary 게시물 상세
// @Description 게시물 상세
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId} [get]
func GetPost(c *gin.Context) {
}
