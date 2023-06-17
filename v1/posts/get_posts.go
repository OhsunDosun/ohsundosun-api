package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

type post struct {
	ID           uint       `json:"-"`
	UUID         model.UUID `json:"uuid"  binding:"required" example:"test"`
	MBTI         string     `json:"mbti" binding:"required" example:"INTP"`
	Type         string     `json:"type"  binding:"required" example:"DAILY"`
	Nickname     string     `json:"nickname"  binding:"required" example:"test"`
	Title        string     `json:"title"  binding:"required" example:"test"`
	Content      string     `json:"content"  binding:"required" example:"test"`
	Images       string     `json:"images"  binding:"required" example:"test.png,test.png"`
	CreatedAt    time.Time  `json:"createdAt"  binding:"required"`
	LikeCount    int8       `json:"likeCount" binding:"required" example:"0"`
	CommentCount int8       `json:"commentCount" binding:"required" example:"0"`
	IsLike       bool       `json:"isLike" binding:"required"`
	IsMine       bool       `json:"isMine" binding:"required"`
}

// GetPosts godoc
// @Tags Posts
// @Summary 게시물 리스트
// @Description 게시물 리스트
// @Security AppAuth
// @Param request query posts.GetPosts.request true "query params"
// @Success 200 {object} model.DataResponse{data=posts.GetPosts.data} "success"
// @Success 404 {object} model.DefaultResponse "not_found_posts"
// @Router /v1/posts [get]
func GetPosts(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Sort    string  `form:"sort" enums:"NEW,LIKE" binding:"required" example:"NEW"`
		Keyword *string `form:"keyword"`
		LastKey *uint   `form:"lastKey"`
		Limit   *int    `form:"limit"`
		MBTI    *string `form:"mbti" enums:"INTJ,INTP,ENTJ,ENTP,INFJ,INFP,ENFJ,ENFP,ISFJ,ISTJ,ESFJ,ESTJ,ISFP,ISTP,ESFP,ESTP"`
		Type    *string `form:"type" enums:"DAILY,LOVE,FRIEND"`
	}

	type data struct {
		List    []post `json:"list" binding:"required"`
		LastKey *uint  `json:"lastKey"`
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

	var posts []post

	postsSelect := db.DB.Model(&model.Post{})
	postsSelect = postsSelect.Select("posts.id, posts.uuid, posts.mbti, posts.type, users.nickname, posts.title, posts.content, posts.images, posts.created_at, posts.like_count, posts.comment_count, COUNT(post_likes.id) > 0 as is_like, posts.user_id = ? as is_mine", user.ID)
	postsSelect = postsSelect.Joins("left join users on posts.user_id = users.id")
	postsSelect = postsSelect.Joins("left join post_likes on posts.id = post_likes.post_id and post_likes.user_id = ?", user.ID)

	postsSelect = postsSelect.Where("posts.active", true)

	if req.MBTI != nil && *req.MBTI != "" {
		postsSelect = postsSelect.Where("posts.mbti", *req.MBTI)
	}
	if req.Type != nil && *req.Type != "" {
		postsSelect = postsSelect.Where("posts.type", *req.Type)
	}

	if req.Keyword != nil && *req.Keyword != "" {
		postsSelect = postsSelect.Where("posts.title LIKE ?", "%"+*req.Keyword+"%")
	}

	if req.LastKey != nil && *req.LastKey != 0 {
		postsSelect = postsSelect.Where("posts.id < ?", *req.LastKey)
	}

	if req.Sort == "NEW" {
		postsSelect = postsSelect.Order("posts.created_at desc")
	} else {
		postsSelect = postsSelect.Order("posts.like_count desc")
	}

	if req.Limit != nil && *req.Limit != 0 {
		postsSelect = postsSelect.Limit(*req.Limit)
	}

	postsSelect = postsSelect.Group("posts.id")

	postsSelect.Find(&posts)

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_posts",
		})
		c.Abort()
		return
	}

	if len(posts) == *req.Limit {
		c.JSON(http.StatusOK, &model.DataResponse{
			Message: "success",
			Data: &data{
				List:    posts,
				LastKey: &posts[*req.Limit-1].ID,
			},
		})
	} else {
		c.JSON(http.StatusOK, &model.DataResponse{
			Message: "success",
			Data: &data{
				List: posts,
			},
		})
	}
}
