package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// UpdatePost godoc
// @Tags Posts
// @Summary 게시글 수정
// @Description 게시글 수정
// @Security AppAuth
// @Param request body posts.UpdatePost.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId} [put]
func UpdatePost(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	type request struct {
		Title   string   `json:"title" binding:"required,max=30" example:"test"`
		Content string   `json:"content" binding:"required,max=6000" example:"test"`
		Type    string   `json:"type" enums:"DAILY,LOVE,FRIEND" binding:"required" example:"DAILY"`
		Images  []string `json:"images" binding:"required" exmaple:"[]"`
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

	if post.UserID != user.ID {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	if err := db.DB.Model(&post).Updates(model.Post{
		Type:    enum.PostType(req.Type),
		Title:   req.Title,
		Content: req.Content,
		Images:  strings.Join(req.Images, ","),
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
