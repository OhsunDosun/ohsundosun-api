package posts

import "github.com/gin-gonic/gin"

// UpdatePost godoc
// @Tags Posts
// @Summary 게시물 수정
// @Description 게시물 수정
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId} [put]
func UpdatePost(c *gin.Context) {
}
