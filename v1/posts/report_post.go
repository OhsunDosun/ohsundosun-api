package posts

import "github.com/gin-gonic/gin"

// ReportPost godoc
// @Tags Posts
// @Summary 게시물 신고
// @Description 게시물 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId}/report [post]
func ReportPost(c *gin.Context) {
}
