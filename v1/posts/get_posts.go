package posts

import (
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"strings"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
)

// GetPosts godoc
// @Tags Posts
// @Summary 게시물 리스트
// @Description 게시물 리스트
// @Security AppAuth
// @Param request query posts.GetPosts.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]posts.GetPosts.data} "success"
// @Success 404 {object} model.DefaultResponse "not_found_posts"
// @Router /v1/posts [get]
func GetPosts(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Sort    string  `form:"sort" enums:"NEW,LIKE" binding:"required" example:"NEW"`
		Keyword *string `form:"keyword"`
		LastKey *string `form:"lastKey"`
		Limit   *int    `form:"limit"`
		MBTI    *string `form:"mbti" enums:"INTJ,INTP,ENTJ,ENTP,INFJ,INFP,ENFJ,ENFP,ISFJ,ISTJ,ESFJ,ESTJ,ISFP,ISTP,ESFP,ESTP"`
		Type    *string `form:"type" enums:"DAILY,LOVE,FRIEND"`
	}

	type data struct {
		Key          string   `json:"key"  binding:"required" example:"test"`
		UserKey      string   `json:"userKey"  binding:"required" example:"test"`
		MBTI         string   `json:"mbti" binding:"required" example:"INTP"`
		Type         string   `json:"type"  binding:"required" example:"DAILY"`
		Nickname     string   `json:"nickname"  binding:"required" example:"test"`
		Title        string   `json:"title"  binding:"required" example:"test"`
		Content      string   `json:"content"  binding:"required" example:"test"`
		Images       []string `json:"images"  binding:"required" example:"test.png,test.png"`
		CreatedAt    int64    `json:"createdAt"  binding:"required"`
		LikeCount    int8     `json:"likeCount" binding:"required" example:"0"`
		CommentCount int8     `json:"commentCount" binding:"required" example:"0"`
		IsLike       bool     `json:"isLike" binding:"required"`
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
	if req.MBTI != nil && *req.MBTI != "" {
		queryData["mbti"] = enum.StringToMBTI(strings.ToUpper(*req.MBTI))
	}
	if req.Type != nil && *req.Type != "" {
		queryData["type"] = enum.StringToPostType(strings.ToUpper(*req.Type))
	}

	query := base.Query{queryData}

	if req.Keyword != nil && *req.Keyword != "" {
		queryTitleData := make(map[string]interface{})
		for key, value := range queryData {
			queryTitleData[key] = value
		}
		queryTitleData["title?contains"] = req.Keyword

		queryContentData := make(map[string]interface{})
		for key, value := range queryData {
			queryContentData[key] = value
		}
		queryContentData["content?contains"] = req.Keyword

		query = base.Query{queryTitleData, queryContentData}
	}

	list := []*data{}

	if req.Sort == "NEW" {
		var result []*model.Post

		deta.BasePost.Fetch(&base.FetchInput{
			Q:       query,
			Dest:    &result,
			Limit:   *req.Limit,
			LastKey: *req.LastKey,
		})

		if len(result) == 0 {
			c.JSON(http.StatusNotFound, &model.DefaultResponse{
				Message: "not_found_posts",
			})
			c.Abort()
			return
		}

		for _, post := range result {
			queryData := make(map[string]interface{})
			queryData["postKey"] = post.Key
			queryData["userKey"] = user.Key

			query := base.Query{queryData}

			var result []*model.PostLike

			deta.BasePostLike.Fetch(&base.FetchInput{
				Q:    query,
				Dest: &result,
			})

			list = append(list, &data{
				Key:          post.Key,
				UserKey:      post.UserKey,
				MBTI:         post.MBTI.String(),
				Type:         post.Type.String(),
				Nickname:     post.Nickname,
				Title:        post.Title,
				Content:      post.Content,
				Images:       post.Images,
				CreatedAt:    post.CreatedAt,
				LikeCount:    post.LikeCount,
				CommentCount: post.CommentCount,
				IsLike:       len(result) > 0,
			},
			)
		}
	} else {
		var result []*model.LikeSortPost

		deta.BaseLikeSortPost.Fetch(&base.FetchInput{
			Q:       query,
			Dest:    &result,
			Limit:   *req.Limit,
			LastKey: *req.LastKey,
		})

		if len(result) == 0 {
			c.JSON(http.StatusNotFound, &model.DefaultResponse{
				Message: "not_found_posts",
			})
			c.Abort()
			return
		}

		for _, post := range result {
			queryData := make(map[string]interface{})
			queryData["postKey"] = post.Key
			queryData["userKey"] = user.Key

			query := base.Query{queryData}

			var result []*model.PostLike

			deta.BasePostLike.Fetch(&base.FetchInput{
				Q:    query,
				Dest: &result,
			})

			list = append(list, &data{
				Key:          post.PostKey,
				UserKey:      post.UserKey,
				MBTI:         post.MBTI.String(),
				Type:         post.Type.String(),
				Nickname:     post.Nickname,
				Title:        post.Title,
				Content:      post.Content,
				Images:       post.Images,
				CreatedAt:    post.CreatedAt,
				LikeCount:    post.LikeCount,
				CommentCount: post.CommentCount,
				IsLike:       len(result) > 0,
			},
			)
		}
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    &list,
	})
}
