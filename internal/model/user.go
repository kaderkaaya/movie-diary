package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string      `gorm:"uniqueIndex;size:32;not null"`
	Email        string      `gorm:"uniqueIndex;size:128;not null"`
	PasswordHash string      `gorm:"size:128;not null"`
	Movies       []UserMovie `gorm:"foreignKey:UserID"`
	Quotes       []Quote     `gorm:"foreignKey:UserID"`
}
