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
// @Summary 게시물 답글 삭제
// @Description 게시물 답글 삭제
// @Security AppAuth
// @Success 201 {object} model.DefaultResponse "success"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_comment_reply"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId}/comments/{commentId}/reply/{replyId} [delete]
func DeleteCommentReply(c *gin.Context) {
	postId := c.Param("postId")
	commentId := c.Param("commentId")
	replyId := c.Param("replyId")

	user := c.MustGet("user").(model.User)

	var comment model.Comment

	err := deta.BaseComment.Get(commentId, &comment)
	if err != nil || comment.PostKey != postId {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comment_reply",
		})
		c.Abort()
		return
	}

	var reply *model.Reply
	newReplys := []model.Reply{}

	for _, v := range comment.Replys {
		if v.Key == replyId {
			reply = &v
			v.Active = false
			v.InActiveAt = sql.NullInt64{
				Int64: time.Now().Unix(),
				Valid: true,
			}
		}
		newReplys = append(newReplys, v)
	}

	if reply == nil {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comment_reply",
		})
		c.Abort()
		return
	}

	if reply.UserKey != user.Key {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	updatesComment := base.Updates{
		"replys": newReplys,
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
		"commentCount": deta.BasePost.Util.Increment(-1),
	}

	deta.BasePost.Update(postId, updatesPost)

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
