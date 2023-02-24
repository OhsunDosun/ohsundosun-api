package model

import (
	"ohsundosun-api/enum"
)

type User struct {
	Key       string    `json:"key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	MBTI      enum.MBTI `json:"mbti"`
	CreatedAt int64     `json:"createdAt"`
}
