package http

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	handlers "moviediary/internal/http/handlers"
	utils "moviediary/pkg/utils"
)

func MovieDiaryRouter(authHandler *handlers.AuthHandler, tokenHandler *handlers.TokenHandler, movieHandler *handlers.MovieHandler) *gin.Engine {
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
	//movie.POST("/add-movie", movieHandler.AddMovie)
	movie.GET("/movie-detail", movieHandler.GetMovieDetail)
	//movie.POST("/delete-movie", movieHandler.DeleteMovie)
	//

	return r
}
