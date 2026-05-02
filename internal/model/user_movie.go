package model

import (
	"time"

	"gorm.io/gorm"
)

type UserMovie struct {
	gorm.Model
	UserID     uint      `gorm:"index;not null"`
	MovieID    uint      `gorm:"index;not null"`
	Rating     float32   `gorm:"not null"`
	Comment    string    `gorm:"type:text"`
	WatchedAt  time.Time `gorm:"not null"`
	IsWatched  bool      `gorm:"default:false"`
	IsFavorite bool      `gorm:"default:false"`

	User  User  `gorm:"foreignKey:UserID"`
	Movie Movie `gorm:"foreignKey:MovieID"`
}
