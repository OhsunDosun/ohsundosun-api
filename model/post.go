package model

import (
	"database/sql"
	"ohsundosun-api/enum"
)

type Post struct {
	Key          string        `json:"key"`
	UserKey      string        `json:"userKey"`
	MBTI         enum.MBTI     `json:"mbti"`
	Type         enum.PostType `json:"type"`
	Nickname     string        `json:"nickname"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Images       []string      `json:"images"`
	CreatedAt    int64         `json:"createdAt"`
	Active       bool          `json:"active"`
	LikeCount    int8          `json:"likeCount"`
	CommentCount int8          `json:"commentCount"`

	UpdatedAt  sql.NullInt64 `json:"updatedAt"`
	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}

type PostLike struct {
	Key       string `json:"key"`
	PostKey   string `json:"postKey"`
	UserKey   string `json:"userKey"`
	CreatedAt int64  `json:"createdAt"`
}
