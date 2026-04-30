package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	config "moviediary/internal/config"
	model "moviediary/internal/model"
	model_dto "moviediary/internal/model/dto"
	"net/http"
	"time"
)

type Client struct {
	http   *http.Client
	apiKey string
}

func NewClient(key string) *Client {
	return &Client{
		http:   &http.Client{Timeout: 5 * time.Second},
		apiKey: key,
	}
}

func (c *Client) GetMovies(ctx context.Context, movieType string, genreID int, year int, page int, rating float64) ([]model.Movie, error) {
	var url string
	apiKey := config.Load().TmdbApiKey
	switch movieType {
	case "trending":
		url = fmt.Sprintf(
			"https://api.themoviedb.org/3/trending/movie/day?api_key=%s",
			apiKey,
		)

	case "popular", "top_rated":
		url = fmt.Sprintf(
			"https://api.themoviedb.org/3/movie/%s?api_key=%s&language=en-US&page=%d",
			movieType, apiKey, page,
		)

	default:
		url = fmt.Sprintf(
			"https://api.themoviedb.org/3/discover/movie?api_key=%s&language=en-US&page=%d&include_adult=false&include_video=false&sort_by=popularity.desc&with_genres=%d&primary_release_year=%d&vote_average.gte=%f",
			apiKey, page, genreID, year, rating,
		)
	}

	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var tmdbResp model_dto.TMDBResponse
	err = json.Unmarshal(body, &tmdbResp)
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	for _, tmdbMovie := range tmdbResp.Results {
		var movieYear int
		if len(tmdbMovie.ReleaseDate) >= 4 {
			fmt.Sscanf(tmdbMovie.ReleaseDate[:4], "%d", &movieYear)
		}

		movies = append(movies, model.Movie{
			TmdbID:     tmdbMovie.ID,
			Title:      tmdbMovie.Title,
			Overview:   tmdbMovie.Overview,
			PosterURL:  tmdbMovie.PosterPath,
			Year:       movieYear,
			ImdbRating: float32(tmdbMovie.VoteAverage),
		})
	}

	return movies, nil
}

func (c *Client) SearchMovies(ctx context.Context, movieName string) ([]model.Movie, error) {
	var url string
	apiKey := config.Load().TmdbApiKey
	url = fmt.Sprintf(
		"https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s",
		apiKey, movieName,
	)
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var tmdbResp model_dto.TMDBResponse
	err = json.Unmarshal(body, &tmdbResp)
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	for _, tmdbMovie := range tmdbResp.Results {
		movies = append(movies, model.Movie{
			TmdbID:    tmdbMovie.ID,
			Title:     tmdbMovie.Title,
			Overview:  tmdbMovie.Overview,
			PosterURL: tmdbMovie.PosterPath,
		})
	}
	return movies, nil
}
