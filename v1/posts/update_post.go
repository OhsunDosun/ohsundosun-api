package posts

import (
	"database/sql"
	"mime/multipart"
	"net/http"
	"ohsundosun-api/deta"
	"ohsundosun-api/enum"
	"ohsundosun-api/model"
	"os"
	"strings"
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/deta/deta-go/service/drive"
	"github.com/gin-gonic/gin"
)

// UpdatePost godoc
// @Tags Posts
// @Summary 게시물 수정
// @Description 게시물 수정
// @Security AppAuth
// @Param request formData posts.UpdatePost.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 403 {object} model.DefaultResponse "forbidden"
// @Success 404 {object} model.DefaultResponse "not_found_post"
// @Success 500 {object} model.DefaultResponse "failed_update"
// @Router /v1/posts/{postId} [put]
func UpdatePost(c *gin.Context) {
	postId := c.Param("postId")

	user := c.MustGet("user").(model.User)

	type request struct {
		Title   *string                  `form:"title"  example:"test"`
		Content *string                  `form:"content"  example:"test"`
		Type    *string                  `form:"type" enums:"DAILY,LOVE,FRIEND" example:"DAILY"`
		Images  *[]*multipart.FileHeader `form:"images" swaggerignore:"true"`
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

	if post.UserKey != user.Key {
		c.JSON(http.StatusForbidden, &model.DefaultResponse{
			Message: "forbidden",
		})
		c.Abort()
		return
	}

	updatesPost := base.Updates{
		"updatedAt": sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	}

	if req.Title != nil {
		updatesPost["title"] = req.Title
	}

	if req.Content != nil {
		updatesPost["content"] = req.Content
	}

	if req.Type != nil {
		postType := enum.StringToPostType(strings.ToUpper(*req.Type))
		if postType == 0 {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_request",
			})
			c.Abort()
			return
		}

		updatesPost["type"] = postType
	}

	if req.Images != nil {
		images := []string{}

		for index, image := range *req.Images {
			if image != nil {
				file, err := image.Open()

				if err != nil {
					break
				}

				name, err := deta.DrivePost.Put(&drive.PutInput{
					Name: postId + "/" + image.Filename,
					Body: file,
				})

				if err != nil {
					break
				}

				images = append(images, os.Getenv("APP_HOST")+"/image/post/"+name)
			} else {
				images = append(images, post.Images[index])
			}
		}

		updatesPost["images"] = images
	}

	err = deta.BasePost.Update(postId, updatesPost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}
