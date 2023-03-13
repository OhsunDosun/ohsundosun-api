package model

import (
	"database/sql"
	"ohsundosun-api/enum"
)

type Comment struct {
	Key          string    `json:"key"`
	CommentKey   string    `json:"commentKey"`
	CommentLevel int8      `json:"commentLevel"`
	UserKey      string    `json:"userKey"`
	Nickname     string    `json:"nickname"`
	MBTI         enum.MBTI `json:"mbti"`
	Content      string    `json:"content"`
	CreatedAt    int64     `json:"createdAt"`
	UpdatedAt    int64     `json:"updatedAt"`
	Active       bool      `json:"active"`

	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}
