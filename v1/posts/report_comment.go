package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// ReportComment godoc
// @Tags Posts-Comments
// @Summary 게시글 댓글 신고
// @Description 게시글 댓글 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 404 {object} model.DefaultResponse "not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/comments/{commentId}/report [post]
func ReportComment(c *gin.Context) {
	postId := c.Param("postId")
	commentId := c.Param("commentId")

	user := c.MustGet("user").(model.User)

	var post model.Post

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?)", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	var comment model.Comment

	if err := db.DB.Model(&model.Comment{}).First(&comment, "uuid = UUID_TO_BIN(?)", commentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	report := model.Report{
		Type:     enum.ReportType("Comment"),
		UserID:   user.ID,
		TargetID: comment.ID,
	}

	if err := db.DB.Create(&report).Error; err != nil {
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
