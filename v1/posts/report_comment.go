package posts

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportComment godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 신고
// @Description 게시물 댓글 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 404 {object} model.DefaultResponse "not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/comments/{commentId}/report [post]
func ReportComment(c *gin.Context) {
	postId := c.Param("postId")
	commentId := c.Param("commentId")

	var comment model.Comment

	err := deta.BaseComment.Get(commentId, &comment)
	if err != nil || comment.PostKey != postId {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comment",
		})
		c.Abort()
		return
	}

	report := &model.Report{
		Key:       util.NewULID().String(),
		Type:      enum.COMMENT,
		TargetKey: commentId,
		CreatedAt: time.Now().Unix(),
	}

	_, err = deta.BaseReport.Insert(report)

	if err != nil {
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
