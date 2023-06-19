package model

import (
	"ohsundosun-api/enum"
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	UUID         UUID      `gorm:"uniqueIndex;default:(UUID_TO_BIN(UUID()));not null"`
	Email        string    `gorm:"unique;index;not null"`
	Password     string    `gorm:"not null"`
	Nickname     string    `gorm:"unique;index;not null"`
	MBTI         enum.MBTI `gorm:"type:ENUM('INTJ', 'INTP', 'ENTJ', 'ENTP', 'INFJ', 'INFP', 'ENFJ', 'ENFP', 'ISFJ', 'ISTJ', 'ESFJ', 'ESTJ', 'ISFP', 'ISTP', 'ESFP', 'ESTP');not null"`
	Notification *bool     `gorm:"default:true;not null"`

	Active     *bool      `gorm:"default:true;not null"`
	InActiveAt *time.Time `gorm:"default:null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"default:null"`
}

type UserToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type UserTemporaryPassword struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type UserRating struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	Rating    float32   `gorm:"default:0;not null"`
	Feedback  *string   `gorm:"default:null"`
	CreatedAt time.Time `gorm:"not null"`
}
