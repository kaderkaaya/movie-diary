// internal/model/movie.go
package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	TmdbID     int    `gorm:"uniqueIndex;not null"`
	ImdbID     string `gorm:"index;size:16"`
	Title      string `gorm:"size:255;not null"`
	Overview   string `gorm:"type:text"`
	PosterURL  string `gorm:"size:512"`
	Year       int
	ImdbRating float32
	FetchedAt  time.Time
}
