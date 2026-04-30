package model

type ListMoviesRequest struct {
	MovieType string  `uri:"movie_type" form:"movie_type" binding:"omitempty,oneof=trending popular top_rated discover"`
	GenreID   int     `form:"genre_id" binding:"omitempty"`
	Year      int     `form:"year" binding:"omitempty,min=1900,max=2100"`
	Page      int     `form:"page" binding:"omitempty,min=1"`
	Rating    float64 `form:"rating" binding:"omitempty,min=0,max=10"`
}

type MovieListItemResponse struct {
	TmdbID      int     `json:"tmdb_id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	PosterURL   string  `json:"poster_url"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
}

type MovieListResponse struct {
	Page       int                     `json:"page"`
	TotalPages int                     `json:"total_pages"`
	TotalItems int                     `json:"total_items"`
	Items      []MovieListItemResponse `json:"items"`
}

type MovieDetailResponse struct {
	TmdbID      int     `json:"tmdb_id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	PosterURL   string  `json:"poster_url"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
}

type TMDBMovie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	PosterPath  string  `json:"poster_path"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
}

type TMDBResponse struct {
	Page         int         `json:"page"`
	Results      []TMDBMovie `json:"results"`
	TotalPages   int         `json:"total_pages"`
	TotalResults int         `json:"total_results"`
}

type SearchMoviesRequest struct {
	MovieName string `form:"movie_name" binding:"omitempty,min=3,max=255"`
}
