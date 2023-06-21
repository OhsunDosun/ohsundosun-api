package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
)

// GetPost godoc
// @Tags Posts
// @Summary 게시글 상세
// @Description 게시글 상세
// @Security AppAuth
// @Success 200 {object} model.DataResponse{data=posts.post} "success"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Router /v1/posts/{postId} [get]
func GetPost(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	var data post

	postsSelect := db.DB.Model(&model.Post{})
	postsSelect = postsSelect.Select("posts.id, posts.uuid, users.uuid as user_uuid, posts.mbti, posts.type, users.nickname, posts.title, posts.content, posts.images, posts.created_at, posts.like_count, posts.comment_count, COUNT(post_likes.id) > 0 as is_like, posts.user_id = ? as is_mine", user.ID)
	postsSelect = postsSelect.Joins("left join users on posts.user_id = users.id")
	postsSelect = postsSelect.Joins("left join post_likes on posts.id = post_likes.post_id and post_likes.user_id = ?", user.ID)
	postsSelect = postsSelect.Where("posts.active", true)
	postsSelect = postsSelect.Where("posts.uuid = UUID_TO_BIN(?)", postId)
	postsSelect = postsSelect.Group("posts.id")

	if err := postsSelect.First(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    data,
	})
}
