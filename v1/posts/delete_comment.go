package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeletePost godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 삭제
// @Description 게시물 댓글 삭제
// @Security AppAuth
// @Success 200 {object} model.DefaultResponse "success"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_post, not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId}/comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
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
			Message: "not_found_comment",
		})
		c.Abort()
		return
	}

	if comment.UserID != user.ID {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	var comments []model.Comment

	active := true

	db.DB.Model(&model.Comment{}).Where(model.Comment{
		GroupID: &comment.ID,
		Active:  &active,
	}).Where("id > ?", comment.ID).Where("level > ?", comment.Level).Find(&comments)

	active = false
	now := time.Now()

	if err := db.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&comment).Updates(model.Comment{
			Active:     &active,
			InActiveAt: &now,
		}).Error; err != nil {
			return err
		}

		for _, comment := range comments {
			if err := tx.Model(comment).Updates(model.Comment{
				Active:     &active,
				InActiveAt: &now,
			}).Error; err != nil {
				return err
			}
		}

		if err := tx.Exec("UPDATE posts SET comment_count = comment_count + ? WHERE id = ?", -(len(comments) + 1), post.ID).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
