package posts

import (
	"fmt"
	"net/http"
	"ohsundosun-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetComments godoc
// @Tags Posts-Comments
// @Summary 게시물 댓글 리스트
// @Description 게시물 댓글 리스트
// @Security AppAuth
// @Param   limit       query     int        false  "Limit"
// @Param   lastKey     query     string     false  "LastKey"
// @Success 200 {object} model.DefaultResponse "success"
// @Router /v1/posts/{postId}/comments [get]
func GetComments(c *gin.Context) {

	limitString := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}
	fmt.Println(limit)
}
