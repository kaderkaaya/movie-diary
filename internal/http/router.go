package http

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	handlers "moviediary/internal/http/handlers"
	utils "moviediary/pkg/utils"
)

func MovieDiaryRouter(authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //r.Use olusturunca middle olsuturduk.
	r.Use(gin.Recovery())
	r.Use(timeout.New(
		timeout.WithTimeout(10*time.Second), //burda eğer bir endpointin süresi verdiğimiz süreden fazla sürerse timeout hatasi verir.
		timeout.WithResponse(func(c *gin.Context) {
			utils.Fail(c, http.StatusRequestTimeout, "Request timeout")
		}),
	))

	auth := r.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	return r
}
