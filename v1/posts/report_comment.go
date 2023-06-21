package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"time"

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

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?) AND active = true", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	var comment model.Comment

	if err := db.DB.Model(&model.Comment{}).First(&comment, "uuid = UUID_TO_BIN(?) AND active = true", commentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	var reports []model.Report

	db.DB.Model(&model.Report{}).Where(model.Report{
		Type:     enum.ReportType("Comment"),
		UserID:   user.ID,
		TargetID: comment.ID,
	}).Find(&reports)

	if len(reports) == 0 {
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

		db.DB.Model(&model.Report{}).Where(model.Report{
			Type:     enum.ReportType("Comment"),
			TargetID: comment.ID,
		}).Find(&reports)

		if len(reports) > 2 {
			var comments []model.Comment

			active := true

			db.DB.Model(&model.Comment{}).Where(model.Comment{
				GroupID: &comment.ID,
				Active:  &active,
			}).Where("id > ?", comment.ID).Where("level > ?", comment.Level).Find(&comments)

			active = false
			now := time.Now()

			db.DB.Model(&comment).Updates(model.Comment{
				Active:     &active,
				InActiveAt: &now,
			})

			for _, comment := range comments {
				db.DB.Model(comment).Updates(model.Comment{
					Active:     &active,
					InActiveAt: &now,
				})
			}

			db.DB.Exec("UPDATE posts SET comment_count = comment_count + ? WHERE id = ?", -(len(comments) + 1), post.ID)
		}
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
