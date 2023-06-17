package posts

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

type comment struct {
	ID        uint       `json:"-"`
	UUID      model.UUID `json:"uuid"  binding:"required" example:"test"`
	MBTI      string     `json:"mbti" binding:"required" example:"INTP"`
	Nickname  string     `json:"nickname"  binding:"required" example:"test"`
	Level     uint       `json:"level"  binding:"required" example:"0"`
	Content   string     `json:"content"  binding:"required" example:"test"`
	CreatedAt time.Time  `json:"createdAt"  binding:"required"`
	IsMine    bool       `json:"isMine" binding:"required"`
}

// GetComments godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 리스트
// @Description 게시물 댓글 리스트
// @Security AppAuth
// @Param request query posts.GetComments.request true "query params"
// @Success 200 {object} model.DataResponse{data=posts.GetComments.data} "success"
// @Success 404 {object} model.DefaultResponse "not_found_comments"
// @Router /v1/posts/{postId}/comments [get]
func GetComments(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	type request struct {
		LastKey *int `form:"lastKey"`
		Limit   *int `form:"limit"`
	}

	type data struct {
		List    []comment `json:"list" binding:"required"`
		LastKey *int      `json:"lastKey"`
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

	var post model.Post

	if err := db.DB.Model(&model.Post{}).First(&post, "uuid = UUID_TO_BIN(?)", postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "not_found_post",
		})
		c.Abort()
		return
	}

	var comments []comment

	commentsSelect := db.DB.Model(&model.Comment{})
	commentsSelect = commentsSelect.Select("comments.id, comments.uuid, users.mbti, users.nickname, comments.level, comments.content, comments.created_at, comments.user_id = ? as is_mine", user.ID)
	commentsSelect = commentsSelect.Joins("left join users on comments.user_id = users.id")

	commentsSelect = commentsSelect.Where("comments.post_id", post.ID)
	commentsSelect = commentsSelect.Where("comments.active", true)

	commentsSelect = commentsSelect.Order("comments.group_id desc, comments.level asc")

	if req.Limit != nil && *req.Limit != 0 && req.LastKey != nil {
		commentsSelect = commentsSelect.Limit(*req.Limit).Offset(*req.LastKey)
	}

	commentsSelect.Find(&comments)

	if len(comments) == 0 {
		c.JSON(http.StatusNotFound, &model.DefaultResponse{
			Message: "not_found_comments",
		})
		c.Abort()
		return
	}

	lastKey := 0

	if req.LastKey != nil && *req.LastKey != 0 {
		lastKey = *req.LastKey
	}

	if req.Limit != nil && *req.Limit != 0 {
		lastKey += *req.Limit
	}

	if len(comments) == *req.Limit {
		c.JSON(http.StatusOK, &model.DataResponse{
			Message: "success",
			Data: &data{
				List:    comments,
				LastKey: &lastKey,
			},
		})
	} else {
		c.JSON(http.StatusOK, &model.DataResponse{
			Message: "success",
			Data: &data{
				List: comments,
			},
		})
	}
}
