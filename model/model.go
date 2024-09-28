package model

import (
	"time"
)

type Blog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Category  string    `gorm:"type:varchar(100)"`
	Tags      string    `gorm:"type:varchar(255)"`
	ViewCount int       `gorm:"default:0"`
	Author    string    `gorm:"type:varchar(255)"`
}
