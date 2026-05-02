package repository

import (
	"context"
	"errors"
	model "moviediary/internal/model"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (movieRepository *MovieRepository) GetByTmdbID(ctx context.Context, tmdbID int) (*model.Movie, error) {
	var movie model.Movie
	if err := movieRepository.db.WithContext(ctx).Where("tmdb_id = ?", tmdbID).First(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &movie, nil
}

func (movieRepository *MovieRepository) Create(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	if err := movieRepository.db.WithContext(ctx).Create(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}
