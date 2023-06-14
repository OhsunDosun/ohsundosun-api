package util

import (
	"ohsundosun-api/db"
	"ohsundosun-api/model"
)

func VerifyEmail(email *string) bool {
	var user *model.User

	result := db.DB.Model(&model.User{}).First(&user, "email = ?", email)

	return result.RowsAffected == 0
}

func VerifyNickname(nickname *string) bool {
	var user *model.User

	result := db.DB.Model(&model.User{}).First(&user, "nickname = ?", nickname)

	return result.RowsAffected == 0
}
