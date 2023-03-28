package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/deta/deta-go/service/base"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	deta "ohsundosun-api/deta"
	docs "ohsundosun-api/docs"
	"ohsundosun-api/model"
	v1 "ohsundosun-api/v1"
)

func setEnv() {
	if os.Getenv("APP_MODE") != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setSwagger(r *gin.Engine) {
	if os.Getenv("APP_MODE") != "prod" {
		docs.SwaggerInfo.Version = "0.0.1"
		docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Title = "오순도순 API"

		r.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{
			os.Getenv("SWAGGER_ID"): os.Getenv("SWAGGER_PWD"),
		}), ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func Actions(c *gin.Context) {
	type event struct {
		Id      string `json:"id" binding:"required" example:"test"`
		Trigger string `json:"trigger" binding:"required" example:"schedule"`
	}

	type request struct {
		Event event `json:"event" binding:"required"`
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

	switch req.Event.Id {
	case "sort_post":
		var oldPost []*model.LikeSortPost

		deta.BaseLikeSortPost.Fetch(&base.FetchInput{
			Q:    base.Query{},
			Dest: &oldPost,
		})

		for _, post := range oldPost {
			deta.BaseLikeSortPost.Delete(post.Key)
		}

		var posts []*model.Post

		deta.BasePost.Fetch(&base.FetchInput{
			Q:    base.Query{},
			Dest: &posts,
		})

		newPost := []*model.LikeSortPost{}

		for _, post := range posts {
			newPost = append(newPost, &model.LikeSortPost{
				PostKey:      post.Key,
				UserKey:      post.UserKey,
				MBTI:         post.MBTI,
				Type:         post.Type,
				Nickname:     post.Nickname,
				Title:        post.Title,
				Content:      post.Content,
				Images:       post.Images,
				CreatedAt:    post.CreatedAt,
				Active:       post.Active,
				LikeCount:    post.LikeCount,
				CommentCount: post.CommentCount,
			},
			)
		}

		sort.Slice(newPost, func(i, j int) bool {
			return newPost[i].LikeCount > newPost[j].LikeCount
		})

		for index, post := range newPost {
			deta.BaseLikeSortPost.Insert(
				&model.LikeSortPost{
					Key:          strconv.Itoa(index),
					PostKey:      post.PostKey,
					UserKey:      post.UserKey,
					MBTI:         post.MBTI,
					Type:         post.Type,
					Nickname:     post.Nickname,
					Title:        post.Title,
					Content:      post.Content,
					Images:       post.Images,
					CreatedAt:    post.CreatedAt,
					Active:       post.Active,
					LikeCount:    post.LikeCount,
					CommentCount: post.CommentCount,
				},
			)
		}

	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
}

// @securityDefinitions.apikey AppAuth
// @in header
// @name App-Key

func main() {
	setEnv()

	r := gin.Default()

	setSwagger(r)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("APP_NAME"))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/image/post/*name", func(c *gin.Context) {
		name := strings.Replace(c.Param("name"), "/", "", 1)

		file, err := deta.DrivePost.Get(name)
		if err != nil {
			fmt.Println("failed to get file with name:", name)
			return
		}
		defer file.Close()

		image, err := io.ReadAll(file)

		if err != nil {
			fmt.Println("failed to readall file")
			return
		}
		mimeType := http.DetectContentType(image)

		c.Data(http.StatusOK, mimeType, image)
	})

	r.POST("/__space/v0/actions", Actions)

	v1.SetRoute(r)

	r.Run()
}
