package posts

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/posts")
	{

		// post
		auth.GET("", middleware.CheckAccessToken(), GetPosts)
		auth.GET(":postId", middleware.CheckAccessToken(), GetPost)
		auth.POST("", middleware.CheckAccessToken(), AddPost)
		auth.PUT(":postId", middleware.CheckAccessToken(), UpdatePost)
		auth.DELETE(":postId", middleware.CheckAccessToken(), DeletePost)
		auth.POST(":postId/report", middleware.CheckAccessToken(), ReportPost)
		auth.PATCH(":postId/like", middleware.CheckAccessToken(), UpdateLike)

		// comment
		auth.GET(":postId/comments", middleware.CheckAccessToken(), GetComments)
		auth.POST(":postId/comments", middleware.CheckAccessToken(), AddComment)
		auth.DELETE(":postId/comments/:commentId", middleware.CheckAccessToken(), DeleteComment)
		auth.POST(":postId/comments/:commentId/report", middleware.CheckAccessToken(), ReportComment)
	}
}
