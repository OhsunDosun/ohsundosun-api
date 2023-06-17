package model

import (
	"ohsundosun-api/enum"
)

type LikeSortPost struct {
	Key          string        `json:"key"`
	PostKey      string        `json:"postKey"`
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
}
