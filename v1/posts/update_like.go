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

	err = deta.BasePost.Get(postId, &post)
	if err != nil || post.Key != postId {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	queryData := make(map[string]interface{})
	queryData["postKey"] = postId
	queryData["userKey"] = user.Key

	query := base.Query{queryData}

	var result []*model.PostLike

	deta.BasePostLike.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	for _, postLike := range result {
		deta.BasePostLike.Delete(postLike.Key)
	}

	if *req.Like {
		p := &model.PostLike{
			Key:       util.NewULID().String(),
			PostKey:   postId,
			UserKey:   user.Key,
			CreatedAt: time.Now().Unix(),
		}

		deta.BasePostLike.Insert(p)

		updatesPost := base.Updates{
			"likeCount": deta.BasePost.Util.Increment(1),
		}

		deta.BasePost.Update(postId, updatesPost)
	} else {
		updatesPost := base.Updates{
			"likeCount": deta.BasePost.Util.Increment(-len(result)),
		}

		deta.BasePost.Update(postId, updatesPost)
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
