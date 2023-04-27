package posts

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// AddCommentReply godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 답글 추가
// @Description 게시물 댓글 답글 추가
// @Security AppAuth
// @Param request body posts.AddCommentReply.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 404 {object} model.DefaultResponse "not_found_comment"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId}/comments/{commentId}/reply [post]
func AddCommentReply(c *gin.Context) {
	commentId := c.Param("commentId")

	user := c.MustGet("user").(model.User)

	type request struct {
		Content string `json:"content" binding:"required,max=6000" example:"test"`
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

	var comment model.Comment

	err = deta.BaseComment.Get(commentId, &comment)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comment",
		})
		c.Abort()
		return
	}

	updates := base.Updates{
		"replys": deta.BaseComment.Util.Append(
			&model.Reply{
				Key:       util.NewULID().String(),
				UserKey:   user.Key,
				Nickname:  user.Nickname,
				MBTI:      user.MBTI,
				Content:   req.Content,
				CreatedAt: time.Now().Unix(),
				Active:    true,
			},
		),
	}

	err = deta.BaseComment.Update(commentId, updates)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
