package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	v1 "ohsundosun-api/v1"

	docs "ohsundosun-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setEnv() {
	if os.Getenv("APP_MODE") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setSwagger(r *gin.Engine) {
	if os.Getenv("APP_MODE") == "dev" {
		docs.SwaggerInfo.Version = "0.0.1"
		docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Title = "오순도순 API"

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

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

	v1.SetRoute(r)

	r.Run()
}
