package posts

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

type dataReply struct {
	Key       string `json:"key"  binding:"required" example:"test"`
	Nickname  string `json:"nickname"  binding:"required" example:"test"`
	MBTI      string `json:"mbti" binding:"required" example:"INTP"`
	Content   string `json:"content"  binding:"required" example:"test"`
	CreatedAt int64  `json:"createdAt"  binding:"required"`
}

// GetComments godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 리스트
// @Description 게시물 댓글 리스트
// @Security AppAuth
// @Param request query posts.GetComments.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]posts.GetComments.data} "success"
// @Success 404 {object} model.DefaultResponse "not_found_comments"
// @Router /v1/posts/{postId}/comments [get]
func GetComments(c *gin.Context) {
	postId := c.Param("postId")

	type request struct {
		LastKey *string `form:"lastKey"`
		Limit   *int    `form:"limit"`
	}

	type data struct {
		Key       string      `json:"key"  binding:"required" example:"test"`
		Nickname  string      `json:"nickname"  binding:"required" example:"test"`
		MBTI      string      `json:"mbti" binding:"required" example:"INTP"`
		Content   string      `json:"content"  binding:"required" example:"test"`
		CreatedAt int64       `json:"createdAt"  binding:"required"`
		Replys    []dataReply `json:"replys" binding:"required"`
	}

	req := &request{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	queryData := make(map[string]interface{})
	queryData["active"] = true
	queryData["postKey"] = postId

	query := base.Query{queryData}

	list := []*data{}

	var result []*model.Comment

	deta.BaseComment.Fetch(&base.FetchInput{
		Q:       query,
		Dest:    &result,
		Limit:   *req.Limit,
		LastKey: *req.LastKey,
	})

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comments",
		})
		c.Abort()
		return
	}

	for _, comment := range result {
		replys := []dataReply{}

		for _, reply := range comment.Replys {
			if reply.Active {
				replys = append(replys, dataReply{
					Key:       reply.Key,
					Nickname:  reply.Nickname,
					MBTI:      reply.MBTI.String(),
					Content:   reply.Content,
					CreatedAt: reply.CreatedAt,
				},
				)
			}
		}

		list = append(list, &data{
			Key:       comment.Key,
			Nickname:  comment.Nickname,
			MBTI:      comment.MBTI.String(),
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			Replys:    replys,
		},
		)
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    &list,
	})
}
