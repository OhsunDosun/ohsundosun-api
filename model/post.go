package model

import (
	"ohsundosun-api/enum"
	"time"
)

type Post struct {
	ID           uint          `gorm:"primaryKey"`
	UUID         UUID          `gorm:"uniqueIndex;default:(UUID_TO_BIN(UUID()));not null"`
	UserID       uint          `gorm:"index;not null"`
	MBTI         enum.MBTI     `gorm:"index;type:ENUM('INTJ', 'INTP', 'ENTJ', 'ENTP', 'INFJ', 'INFP', 'ENFJ', 'ENFP', 'ISFJ', 'ISTJ', 'ESFJ', 'ESTJ', 'ISFP', 'ISTP', 'ESFP', 'ESTP');not null"`
	Type         enum.PostType `gorm:"index;type:ENUM('DAILY', 'LOVE', 'FRIEND');not null"`
	Title        string        `gorm:"index;not null"`
	Content      string        `gorm:"not null"`
	Images       string        `gorm:"not null"`
	LikeCount    int8          `gorm:"index;default:0;not null"`
	CommentCount int8          `gorm:"default:0;not null"`

	Active     *bool      `gorm:"default:true;not null"`
	InActiveAt *time.Time `gorm:"default:null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"default:null"`
}

type PostLike struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	PostID    uint      `gorm:"index;not null"`
	CreatedAt time.Time `gorm:"not null"`
}
