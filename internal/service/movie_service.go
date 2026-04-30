package service

import (
	"context"
	model "moviediary/internal/model/dto"
	provider "moviediary/internal/provider/tmdb"
	repository "moviediary/internal/repository"
	"time"
)

type MovieService struct {
	movieRepository *repository.MovieRepository
	tmdbClient      *provider.Client
}

func NewMovieService(movieRepository *repository.MovieRepository, tmdbClient *provider.Client) *MovieService {
	return &MovieService{movieRepository: movieRepository, tmdbClient: tmdbClient}
}

func (movieService *MovieService) ListMovies(ctx context.Context, movieType string, genreID int, year int, page int, rating float64) (*model.MovieListResponse, error) {

	movies, err := movieService.tmdbClient.GetMovies(ctx, movieType, genreID, year, page, rating)
	if err != nil {
		return nil, err
	}
	totalPages := len(movies)
	totalItems := len(movies)
	items := make([]model.MovieListItemResponse, len(movies))
	for i, movie := range movies {
		items[i] = model.MovieListItemResponse{
			TmdbID:      movie.TmdbID,
			Title:       movie.Title,
			Overview:    movie.Overview,
			PosterURL:   movie.PosterURL,
			ReleaseDate: time.Now().Format("2006-01-02"),
			Rating:      float64(movie.ImdbRating),
		}
	}
	return &model.MovieListResponse{
		Page:       page,
		TotalPages: totalPages,
		TotalItems: totalItems,
		Items:      items,
	}, nil
}
