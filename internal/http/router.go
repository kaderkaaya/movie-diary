package http

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	"moviediary/internal/config"
	handlers "moviediary/internal/http/handlers"
	"moviediary/internal/http/middleware"
	utils "moviediary/pkg/utils"
)

func MovieDiaryRouter(authHandler *handlers.AuthHandler, tokenHandler *handlers.TokenHandler, movieHandler *handlers.MovieHandler, diaryHandler *handlers.DiaryHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //r.Use olusturunca middle olsuturduk.
	r.Use(gin.Recovery())
	r.Use(timeout.New(
		timeout.WithTimeout(10*time.Second), //burda eğer bir endpointin süresi verdiğimiz süreden fazla sürerse timeout hatasi verir.
		timeout.WithResponse(func(c *gin.Context) {
			utils.Fail(c, http.StatusRequestTimeout, "Request timeout")
		}),
	))
	//Auth
	auth := r.Group("/auth")
	//auth.Use(middleware.AuthMiddleware(config.Load().JwtSecret))
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	//Token
	token := r.Group("/token")
	token.POST("/refresh", tokenHandler.RefreshToken)

	//Movie
	movie := r.Group("/movie")
	movie.GET("/list-movies/:movie_type", movieHandler.ListMovies)
	movie.GET("/search-movies", movieHandler.SearchMovies)
	movie.GET("/movie-detail", movieHandler.GetMovieDetail)

	//Diary
	diary := r.Group("/diary")
	diary.Use(middleware.AuthMiddleware(config.Load().JwtSecret))
	diary.POST("/add-diary", diaryHandler.AddDiary)
	diary.POST("/remove-diary", diaryHandler.RemoveDiary)
	//diary.GET("/get-diary-list", diaryHandler.GetDiaryList)

	return r
}
