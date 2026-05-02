package repository

import (
	"context"
	"errors"
	model "moviediary/internal/model"
	"time"

	"gorm.io/gorm"
)

type DiaryRepository struct {
	db *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) *DiaryRepository {
	return &DiaryRepository{db: db}
}

func (diaryRepository *DiaryRepository) AddDiary(ctx context.Context, userId uint, movieId int, comment string, rating float64, watchedAt time.Time) (*model.UserMovie, error) {
	diary := &model.UserMovie{
		UserID:    userId,
		MovieID:   uint(movieId),
		Comment:   comment,
		Rating:    float32(rating),
		WatchedAt: watchedAt,
		IsWatched: true,
	}
	if err := diaryRepository.db.WithContext(ctx).Create(diary).Error; err != nil {
		return nil, err
	}
	return diary, nil
}

func (diaryRepository *DiaryRepository) GetByUserIDAndMovieID(ctx context.Context, userId uint, movieId int) (*model.UserMovie, error) {
	var diary model.UserMovie
	if err := diaryRepository.db.WithContext(ctx).Where("user_id = ? AND movie_id = ?", userId, movieId).First(&diary).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &diary, nil
}
