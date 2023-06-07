package images

import (
	"mime/multipart"
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/model"
	"os"

	"github.com/deta/deta-go/service/drive"
	"github.com/gin-gonic/gin"
)

// AddImage godoc
// @Tags Images
// @Summary 이미지 추가
// @Description 이미지 추가
// @Security AppAuth
// @Param request body images.AddImage.request true "body params"
// @Success 201 {object} model.DataResponse{data=images.AddImage.data} "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Router /v1/images [post]
func AddImage(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Category string                 `form:"category" enums:"POST" binding:"required" example:"POST"`
		Images   []multipart.FileHeader `form:"images" swaggerignore:"true"`
	}

	type data struct {
		Images []string `json:"images" binding:"required"`
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

	images := []string{}

	for _, image := range req.Images {
		file, err := image.Open()

		if err != nil {
			break
		}

		name, err := deta.DrivePost.Put(&drive.PutInput{
			Name: user.Key + "/" + image.Filename,
			Body: file,
		})

		if err != nil {
			break
		}

		images = append(images, os.Getenv("APP_HOST")+"/image/post/"+name)
	}

	c.JSON(http.StatusCreated, &model.DataResponse{
		Message: "success",
		Data: &data{
			Images: images,
		},
	})
}
