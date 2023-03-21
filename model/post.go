package model

import (
	"database/sql"
	"ohsundosun-api/enum"
)

type Post struct {
	Key          string        `json:"key"`
	UserKey      string        `json:"userKey"`
	Nickname     string        `json:"nickname"`
	MBTI         enum.MBTI     `json:"mbti"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Type         enum.PostType `json:"type"`
	Images       []string      `json:"images"`
	CreatedAt    int64         `json:"createdAt"`
	Active       bool          `json:"active"`
	LikeCount    int8          `json:"likeCount"`
	CommentCount int8          `json:"commentCount"`

	UpdatedAt  sql.NullInt64 `json:"updatedAt"`
	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}
