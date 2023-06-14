package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateLike godoc
// @Tags Posts
// @Summary 좋아요
// @Description 좋아요
// @Security AppAuth
// @Param request body posts.UpdateLike.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId}/like [patch]
func UpdateLike(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	type request struct {
		Like *bool `json:"like" binding:"required"`
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

	var postLikes []model.PostLike

	db.DB.Model(&model.PostLike{}).Where(model.PostLike{
		UserID: user.ID,
		PostID: post.ID,
	}).Find(&postLikes)

	db.DB.Transaction(func(tx *gorm.DB) error {
		if *req.Like {
			if len(postLikes) == 0 {
				postLike := model.PostLike{
					UserID: user.ID,
					PostID: post.ID,
				}

				if err := tx.Create(&postLike).Error; err != nil {
					return err
				}

				if err := tx.Exec("UPDATE posts SET like_count = like_count + 1 WHERE id = ?", post.ID).Error; err != nil {
					return err
				}
			}
		} else {
			if len(postLikes) != 0 {
				if err := tx.Delete(&postLikes).Error; err != nil {
					return err
				}

				if err := tx.Exec("UPDATE posts SET like_count = like_count + ? WHERE id = ?", -len(postLikes), post.ID).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
