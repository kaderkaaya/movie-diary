package model

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	UserID  uint   `gorm:"index;not null"`
	MovieID uint   `gorm:"index;not null"`
	Content string `gorm:"type:text;not null"`
	Likes   int    `gorm:"default:0"`

	User  User  `gorm:"foreignKey:UserID"`
	Movie Movie `gorm:"foreignKey:MovieID"`
}
