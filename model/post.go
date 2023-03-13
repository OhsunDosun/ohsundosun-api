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
	UpdatedAt    int64         `json:"updatedAt"`
	Active       bool          `json:"active"`
	LikeCount    int8          `json:"likeCount"`
	CommentCount int8          `json:"commentCount"`

	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}
