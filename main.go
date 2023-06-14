package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ohsundosun-api/db"

	docs "ohsundosun-api/docs"
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
		host := strings.Replace(os.Getenv("APP_HOST"), "https://", "", 1)
		host = strings.Replace(host, "http://", "", 1)

		docs.SwaggerInfo.Version = "0.0.1"
		docs.SwaggerInfo.Host = host
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Title = "오순도순 API"

		r.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{
			os.Getenv("SWAGGER_ID"): os.Getenv("SWAGGER_PWD"),
		}), ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
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

	v1.SetRoute(r)

	r.Run()
}
