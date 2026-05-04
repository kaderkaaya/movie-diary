package model

import (
	"encoding/json"
	"fmt"
	"moviediary/internal/model"
	"strings"
	"time"
)

type FlexibleDateTime struct {
	time.Time
}

func (f *FlexibleDateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return fmt.Errorf("watched_at is required")
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		f.Time = t
		return nil
	}
	if t, err := time.Parse("2006-01-02", s); err == nil {
		f.Time = t.UTC()
		return nil
	}
	return fmt.Errorf("watched_at must be RFC3339 or YYYY-MM-DD, got %q", s)
}

type AddDiaryRequest struct {
	MovieId   int              `json:"tmdb_id" binding:"required,min=1"`
	Comment   string           `json:"comment" binding:"required,min=1,max=255"`
	Rating    float64          `json:"rating" binding:"required,min=0,max=10"`
	WatchedAt FlexibleDateTime `json:"watched_at" binding:"required"`
}

type AddDiaryResponse struct {
	Diary   *model.UserMovie `json:"user_movie"`
	Message string           `json:"message"`
}

type RemoveDiaryRequest struct {
	MovieId int `json:"movie_id" binding:"required,min=1"`
}

type GetDiaryListRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}

type GetDiaryListResponse struct {
	Page       int                `json:"page"`
	Items      []*model.UserMovie `json:"items"`
	TotalPages int                `json:"total_pages"`
	TotalItems int                `json:"total_items"`
	Message    string             `json:"message"`
}
