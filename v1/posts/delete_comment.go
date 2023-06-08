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
// @Tags Posts-Comments
// @Summary 게시물 댓글 삭제
// @Description 게시물 댓글 삭제
// @Security AppAuth
// @Success 200 {object} model.DefaultResponse "success"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId}/comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	postId := c.Param("postId")
	commentId := c.Param("commentId")

	user := c.MustGet("user").(model.User)

	var comment model.Comment

	err := deta.BaseComment.Get(commentId, &comment)
	if err != nil || comment.PostKey != postId {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comment",
		})
		c.Abort()
		return
	}

	if comment.UserKey != user.Key {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	updatesComment := base.Updates{
		"active": false,
		"inActiveAt": sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	}

	err = deta.BaseComment.Update(commentId, updatesComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	updatesPost := base.Updates{
		"commentCount": deta.BasePost.Util.Increment(-1 + -len(comment.Replys)),
	}

	deta.BasePost.Update(postId, updatesPost)

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
