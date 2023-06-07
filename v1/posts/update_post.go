package posts

import (
	"database/sql"
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"strings"
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// UpdatePost godoc
// @Tags Posts
// @Summary 게시물 수정
// @Description 게시물 수정
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
		Title   string   `json:"title" binding:"required,max=30"  example:"test"`
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

	err = deta.BasePost.Get(postId, &post)
	if err != nil || post.Key != postId {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	if post.UserKey != user.Key {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	postType := enum.StringToPostType(strings.ToUpper(req.Type))
	if postType == 0 {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	updatesPost := base.Updates{
		"title":   req.Title,
		"content": req.Content,
		"type":    postType,
		"images":  req.Images,
		"updatedAt": sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	}

	err = deta.BasePost.Update(postId, updatesPost)

	if err != nil {
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
