package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddComment godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 추가
// @Description 게시물 댓글 추가
// @Security AppAuth
// @Param request body posts.AddComment.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 404 {object} model.DefaultResponse "not_found_post, not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/comments [post]
func AddComment(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	type request struct {
		CommentId *string `json:"commentId" example:"test"`
		Content   string  `json:"content" binding:"required,max=6000" example:"test"`
	}

	req := &request{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	var post model.Post

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?)", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	if req.CommentId != nil {
		var comment model.Comment

		if err := db.DB.Model(&model.Comment{}).First(&comment, "uuid = UUID_TO_BIN(?)", req.CommentId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "not_found_comment",
			})
			c.Abort()
			return
		}

		if err := db.DB.Transaction(func(tx *gorm.DB) error {
			comment := model.Comment{
				ParentID: &comment.ID,
				GroupID:  comment.GroupID,
				Level:    comment.Level + 1,
				PostID:   post.ID,
				UserID:   user.ID,
				Content:  req.Content,
			}

			if result := tx.Create(&comment); result.Error != nil {
				return result.Error
			}

			if err := tx.Exec("UPDATE posts SET comment_count = comment_count + 1 WHERE id = ?", post.ID).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "failed_insert",
			})
			c.Abort()
			return
		}
	} else {
		if err := db.DB.Transaction(func(tx *gorm.DB) error {
			comment := model.Comment{
				PostID:  post.ID,
				UserID:  user.ID,
				Content: req.Content,
			}

			if result := tx.Create(&comment); result.Error != nil {
				return result.Error
			}

			if err := tx.Model(&comment).Updates(model.Comment{
				GroupID:  &comment.ID,
				ParentID: &comment.ID,
			}).Error; err != nil {
				return err
			}

			if err := tx.Exec("UPDATE posts SET comment_count = comment_count + 1 WHERE id = ?", post.ID).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
				Message: "failed_insert",
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
