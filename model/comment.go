package model

import (
	"database/sql"
	"ohsundosun-api/enum"
)

type Comment struct {
	Key       string    `json:"key"`
	PostKey   string    `json:"postKey"`
	UserKey   string    `json:"userKey"`
	Nickname  string    `json:"nickname"`
	MBTI      enum.MBTI `json:"mbti"`
	Content   string    `json:"content"`
	CreatedAt int64     `json:"createdAt"`
	Active    bool      `json:"active"`
	Replys    []Reply   `json:"replys"`

	UpdatedAt  sql.NullInt64 `json:"updatedAt"`
	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}

type Reply struct {
	Key       string    `json:"key"`
	UserKey   string    `json:"userKey"`
	Nickname  string    `json:"nickname"`
	MBTI      enum.MBTI `json:"mbti"`
	Content   string    `json:"content"`
	CreatedAt int64     `json:"createdAt"`
	Active    bool      `json:"active"`

	UpdatedAt  sql.NullInt64 `json:"updatedAt"`
	InActiveAt sql.NullInt64 `json:"inActiveAt"`
}
