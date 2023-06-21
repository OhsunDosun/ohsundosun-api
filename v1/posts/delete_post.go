package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

// DeletePost godoc
// @Tags Posts
// @Summary 게시글 삭제
// @Description 게시글 삭제
// @Security AppAuth
// @Success 200 {object} model.DefaultResponse "success"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId} [delete]
func DeletePost(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	var post model.Post

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?)", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	if post.UserID != user.ID {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	active := false
	now := time.Now()

	if err := db.DB.Model(&post).Updates(model.Post{
		Active:     &active,
		InActiveAt: &now,
	}).Error; err != nil {
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
