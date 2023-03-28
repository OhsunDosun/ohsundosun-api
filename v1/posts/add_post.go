package posts

import (
	"mime/multipart"
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"strings"
	"time"

	"github.com/deta/deta-go/service/drive"
	"github.com/gin-gonic/gin"
)

// AddPost godoc
// @Tags Posts
// @Summary 게시물 추가
// @Description 게시물 추가
// @Security AppAuth
// @Param request formData posts.AddPost.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/posts [post]
func AddPost(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Title   string                 `form:"title" binding:"required,max=30" example:"test"`
		Content string                 `form:"content" binding:"required,max=6000" example:"test"`
		Type    string                 `form:"type" enums:"DAILY,LOVE,FRIEND" binding:"required" example:"DAILY"`
		Images  []multipart.FileHeader `form:"images" swaggerignore:"true"`
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

	postType := enum.StringToPostType(strings.ToUpper(req.Type))
	if postType == 0 {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	postKey := util.NewULID().String()

	images := []string{}

	for _, image := range req.Images {
		file, err := image.Open()

		if err != nil {
			break
		}

		name, err := deta.DrivePost.Put(&drive.PutInput{
			Name: postKey + "/" + image.Filename,
			Body: file,
		})

		if err != nil {
			break
		}

		images = append(images, name)
	}

	p := &model.Post{
		Key:       postKey,
		UserKey:   user.Key,
		Nickname:  user.Nickname,
		MBTI:      user.MBTI,
		Title:     req.Title,
		Content:   req.Content,
		Type:      postType,
		Images:    images,
		CreatedAt: time.Now().Unix(),
		Active:    true,
	}

	_, err = deta.BasePost.Insert(p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_insert",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, &model.DefaultResponse{
		Message: "success",
	})
}
