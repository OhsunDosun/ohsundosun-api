package posts

import "github.com/gin-gonic/gin"

// AddComment godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 추가
// @Description 게시물 댓글 추가
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId}/comments [post]
func AddComment(c *gin.Context) {
}
