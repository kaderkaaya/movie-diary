package main

import (
	"log"
	config "moviediary/internal/config"
	apphttp "moviediary/internal/http"
	handlers "moviediary/internal/http/handlers"
	model "moviediary/internal/model"
	"moviediary/internal/provider/tmdb"
	repository "moviediary/internal/repository"
	service "moviediary/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	db, err := model.OpenDB(cfg.DbDsn)
	if err != nil {
		log.Fatal(err)
	}

	err = model.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	//Auth
	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)
	authService := service.NewAuthService(userRepository, tokenRepository)
	authHandler := handlers.NewAuthHandler(authService)
	tokenHandler := handlers.NewTokenHandler(service.NewTokenService(tokenRepository, userRepository))

	//Movie
	movieRepository := repository.NewMovieRepository(db)
	tmdbClient := tmdb.NewClient(cfg.TmdbApiKey)
	movieService := service.NewMovieService(movieRepository, tmdbClient)
	movieHandler := handlers.NewMovieHandler(movieService)

	//Diary
	diaryRepository := repository.NewDiaryRepository(db)
	diaryService := service.NewDiaryService(diaryRepository, movieRepository, tmdbClient)
	diaryHandler := handlers.NewDiaryHandler(diaryService)

	router := apphttp.MovieDiaryRouter(authHandler, tokenHandler, movieHandler, diaryHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server on :%s", port)
	_ = router.Run(":" + port)
}
