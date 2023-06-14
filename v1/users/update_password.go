package users

import (
	"net/http"
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UpdatePaasword godoc
// @Tags Users
// @Summary 비밀번호 변경
// @Description 비밀번호 변경
// @Security AppAuth
// @Param request body users.UpdatePaasword.request true "body params"
// @Success 200 {object} model.DefaultResponse "success"
// @Success 400 {object} model.DefaultResponse "bad_request"
// @Router /v1/users/password [patch]
func UpdatePaasword(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	type request struct {
		Type        string `json:"type" enums:"DEFAULT, NEW_PASSWORD" binding:"required" example:"DEFAULT"`
		OldPassword string `json:"oldPassword" binding:"required,alphanum,min=8,max=16" example:"test1234"`
		NewPassword string `json:"newPassword" binding:"required,alphanum,min=8,max=16" example:"test1234"`
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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_password",
		})
		c.Abort()
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)

	if err := db.DB.Model(&user).Updates(model.User{
		Password: string(hashPassword),
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_update",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DefaultResponse{
		Message: "success",
	})
	// if strings.ToUpper(req.Type) == "NEW_PASSWORD" {
	// 	if !user.NewPasswordCreatedAt.Valid || !user.NewPassword.Valid {
	// 		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
	// 			Message: "bad_password",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	now := time.Now()

	// 	if now.Sub(time.Unix(user.NewPasswordCreatedAt.Int64, 0)) > 3*time.Minute {
	// 		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
	// 			Message: "bad_password",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	err = bcrypt.CompareHashAndPassword([]byte(user.NewPassword.String), []byte(req.OldPassword))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
	// 			Message: "bad_password",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}
	// } else {
	// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
	// 			Message: "bad_password",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}
	// }

	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)

	// err = deta.BaseUser.Update(user.Key, base.Updates{
	// 	"password": string(hashPassword),
	// 	"newPassword": sql.NullString{
	// 		String: "",
	// 		Valid:  false,
	// 	},
	// 	"newPasswordCreatedAt": sql.NullInt64{
	// 		Int64: 0,
	// 		Valid: false,
	// 	},
	// })

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
	// 		Message: "failed_update",
	// 	})
	// 	c.Abort()
	// 	return
	// }

}
