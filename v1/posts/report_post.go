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

// ReportPost godoc
// @Tags Posts
// @Summary 게시물 신고
// @Description 게시물 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/report [post]
func ReportPost(c *gin.Context) {
	postId := c.Param("postId")

	var post model.Post

	err := deta.BasePost.Get(postId, &post)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	report := &model.Report{
		Key:       util.NewULID().String(),
		Type:      enum.POST,
		TargetKey: postId,
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
