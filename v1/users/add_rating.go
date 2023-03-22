package users

import (
	"database/sql"
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"
	"ohsundosun-api/util"
	"time"

	"github.com/gin-gonic/gin"
)

// AddRating godoc
// @Tags Users
// @Summary 평가 추가
// @Description 평가 추가
// @Security AppAuth
// @Param request body users.AddRating.request true "body params"
// @Success 201 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Success 500 {object} model.DefaultResponse "failed_insert"
// @Router /v1/users/rating [post]
func AddRating(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Rating   *float32 `json:"rating" binding:"required" example:"0"`
		Feedback *string  `json:"feedback" example:"test"`
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

	var feedback sql.NullString

	if req.Feedback != nil {
		feedback = sql.NullString{
			String: *req.Feedback,
			Valid:  true,
		}
	} else {
		feedback = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	u := &model.Rating{
		Key:       util.NewULID().String(),
		UserKey:   user.Key,
		Feedback:  feedback,
		CreatedAt: time.Now().Unix(),
	}

	_, err = db.BaseRating.Insert(u)

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
