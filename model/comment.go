package model

import (
	"time"
)

type Comment struct {
	ID       uint   `gorm:"primaryKey"`
	UUID     UUID   `gorm:"uniqueIndex;default:(UUID_TO_BIN(UUID()));not null"`
	GroupID  *uint  `gorm:"index;default:null"`
	ParentID *uint  `gorm:"index;default:null"`
	Level    uint   `gorm:"index;not null;default:0"`
	PostID   uint   `gorm:"index;not null"`
	UserID   uint   `gorm:"index;not null"`
	Content  string `gorm:"not null"`

	Active     *bool      `gorm:"index;default:true;not null"`
	InActiveAt *time.Time `gorm:"default:null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:false;default:null"`
}
