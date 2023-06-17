package model

import (
	"ohsundosun-api/enum"
	"time"
)

type Report struct {
	ID        uint            `gorm:"primaryKey"`
	Type      enum.ReportType `gorm:"index;type:ENUM('POST', 'COMMENT');not null"`
	UserID    uint            `gorm:"index;not null"`
	TargetID  uint            `gorm:"index;not null"`
	CreatedAt time.Time       `gorm:"not null"`
}
