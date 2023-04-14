package model

import (
	"database/sql"
	"ohsundosun-api/enum"
)

type User struct {
	Key          string    `json:"key"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Nickname     string    `json:"nickname"`
	MBTI         enum.MBTI `json:"mbti"`
	CreatedAt    int64     `json:"createdAt"`
	Notification bool      `json:"notification"`
	Active       bool      `json:"active"`
	FCM          []string  `json:"fcm"`

	InActiveAt           sql.NullInt64  `json:"inActiveAt"`
	NewPassword          sql.NullString `json:"newPassword"`
	NewPasswordCreatedAt sql.NullInt64  `json:"newPasswordCreatedAt"`
}
