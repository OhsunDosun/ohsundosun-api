package posts

import "github.com/gin-gonic/gin"

// ReportComment godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 신고
// @Description 게시물 댓글 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId}/comments/{commentId}/report [post]
func ReportComment(c *gin.Context) {
}
