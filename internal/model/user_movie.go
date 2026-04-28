package model

import (
	"time"

	"gorm.io/gorm"
)

type UserMovie struct {
	gorm.Model
	UserID    uint      `gorm:"index;not null"`
	MovieID   uint      `gorm:"index;not null"`
	Rating    float32   `gorm:"not null"`
	Comment   string    `gorm:"type:text"`
	WatchedAt time.Time `gorm:"not null"`

	User  User  `gorm:"foreignKey:UserID"`
	Movie Movie `gorm:"foreignKey:MovieID"`
}
