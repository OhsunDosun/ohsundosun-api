package posts

import (
	"database/sql"
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// DeletePost godoc
// @Tags Posts
// @Summary 게시물 삭제
// @Description 게시물 삭제
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

	err := deta.BasePost.Get(postId, &post)
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

	updatesPost := base.Updates{
		"active": false,
		"inActiveAt": sql.NullInt64{
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
