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

// AddComment godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 추가
// @Description 게시물 댓글 추가
// @Security AppAuth
// @Param request body posts.AddComment.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts/{postId}/comments [post]
func AddComment(c *gin.Context) {
	postId := c.Param("postId")

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

	comment := &model.Comment{
		Key:       util.NewULID().String(),
		PostKey:   postId,
		UserKey:   user.Key,
		Nickname:  user.Nickname,
		MBTI:      user.MBTI,
		Content:   req.Content,
		CreatedAt: time.Now().Unix(),
		Active:    true,
		Replys:    []model.Reply{},
	}

	_, err = deta.BaseComment.Insert(comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_insert",
		})
		c.Abort()
		return
	}

	updatesPost := base.Updates{
		"commentCount": deta.BasePost.Util.Increment(1),
	}

	deta.BasePost.Update(postId, updatesPost)

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
