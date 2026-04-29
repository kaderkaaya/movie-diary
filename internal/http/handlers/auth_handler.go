package handlers

import (
	model_dto "moviediary/internal/model/dto"
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

func (authHandler *AuthHandler) Register(c *gin.Context) {
	var req model_dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := authHandler.service.Register(c.Request.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "user", user, "User logged in successfully")
}

func (authHandler *AuthHandler) Login(c *gin.Context) {
	var req model_dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := authHandler.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "user", user, "User logged in successfully")
}
