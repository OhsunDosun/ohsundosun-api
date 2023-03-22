package model

import (
	"ohsundosun-api/enum"
)

type LikeSortPost struct {
	Key          string        `json:"key"`
	PostKey      string        `json:"postKey"`
	Nickname     string        `json:"nickname"`
	MBTI         enum.MBTI     `json:"mbti"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Type         enum.PostType `json:"type"`
	CreatedAt    int64         `json:"createdAt"`
	LikeCount    int8          `json:"likeCount"`
	CommentCount int8          `json:"commentCount"`
}
