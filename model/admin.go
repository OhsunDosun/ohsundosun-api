package model

import (
	"time"
)

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	UUID     UUID   `gorm:"uniqueIndex;default:(UUID_TO_BIN(UUID()));not null"`
	Email    string `gorm:"unique;index;not null"`
	Password string `gorm:"not null"`
	Nickname string `gorm:"unique;index;not null"`

	Active     *bool      `gorm:"default:true;not null"`
	InActiveAt *time.Time `gorm:"default:null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"default:null"`
}
