package main

import (
	"log"
	"moviediary/internal/model"
	"os"

	"moviediary/internal/config"

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
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server on :%s", port)
	_ = r.Run(":" + port)
}
