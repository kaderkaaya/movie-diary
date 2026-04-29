package handlers

import (
	model "moviediary/internal/model"
	service "moviediary/internal/service"
	utils "moviediary/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// c = request + response + context
// JSON parse → c.ShouldBindJSON
// response → c.JSON
// param → c.Param, c.Query
// context → c.Request.Context()

func (authHandler *AuthHandler) Register(c *gin.Context) {
	//c *gin.Context
	//request’i temsil eder
	//response’u yazmanı sağlar
	//middleware’lerle veri taşır
	//AuthHandler struct’ına bağlı bir method. authHandler.service sayesinde business logic’e erişiyor.
	var req model.RegisterRequest
	//burda gelen json body bu structu doldurucak.
	if err := c.ShouldBindJSON(&req); err != nil {
		// ShouldBindJSON request body’yi al
		//burda gelen hata durumunda fail fonksiyonu çağrılıyor.
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	//gelen dtoları kontrol ediyoruz.
	//Burada Gin, request body’yi RegisterRequest içine map’liyor.
	user, err := authHandler.service.Register(c.Request.Context(), req.Username, req.Email, req.Password)
	// c.Request.Context() bunu service’e geçir
	//gin.Context = request + response + taşıyıcı
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Ok(c, http.StatusOK, user, "User registered successfully")
}
