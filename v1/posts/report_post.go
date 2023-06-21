package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportPost godoc
// @Tags Posts
// @Summary 게시글 신고
// @Description 게시글 신고
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/report [post]
func ReportPost(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	var post model.Post

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?) AND active = true", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	var reports []model.Report

	db.DB.Model(&model.Report{}).Where(model.Report{
		Type:     enum.ReportType("POST"),
		UserID:   user.ID,
		TargetID: post.ID,
	}).Find(&reports)

	if len(reports) == 0 {
		report := model.Report{
			Type:     enum.ReportType("POST"),
			UserID:   user.ID,
			TargetID: post.ID,
		}

		if err := db.DB.Create(&report).Error; err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "failed_insert",
			})
			c.Abort()
			return
		}

		db.DB.Model(&model.Report{}).Where(model.Report{
			Type:     enum.ReportType("POST"),
			TargetID: post.ID,
		}).Find(&reports)

		if len(reports) > 2 {
			active := false
			now := time.Now()

			db.DB.Model(&post).Updates(model.Post{
				Active:     &active,
				InActiveAt: &now,
			})
		}
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
