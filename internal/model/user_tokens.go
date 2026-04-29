package model

import (
	"time"

	"gorm.io/gorm"
)

type UserTokens struct {
	gorm.Model
	UserID    uint      `gorm:"index;not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:UserID"`
}
