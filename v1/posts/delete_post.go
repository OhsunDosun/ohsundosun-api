package posts

import "github.com/gin-gonic/gin"

// DeletePost godoc
// @Tags Posts
// @Summary 게시물 삭제
// @Description 게시물 삭제
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId} [delete]
func DeletePost(c *gin.Context) {
}
