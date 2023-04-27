package posts

import (
	"ohsundosun-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/posts")
	{
		auth.GET("", middleware.CheckAccessToken(), GetPosts)
		auth.GET(":postId", middleware.CheckAccessToken(), GetPost)

		auth.POST("", middleware.CheckAccessToken(), AddPost)
		auth.PUT(":postId", middleware.CheckAccessToken(), UpdatePost)
		auth.DELETE(":postId", middleware.CheckAccessToken(), DeletePost)
		auth.POST(":postId/report", middleware.CheckAccessToken(), ReportPost)

		auth.GET(":postId/comments", middleware.CheckAccessToken(), GetComments)

		auth.POST(":postId/comments", middleware.CheckAccessToken(), AddComment)
		auth.POST(":postId/comments/:commentId/report", middleware.CheckAccessToken(), ReportComment)
		auth.POST(":postId/comments/:commentId/reply", middleware.CheckAccessToken(), AddCommentReply)
	}
}
