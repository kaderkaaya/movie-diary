package handlers

import (
	model_dto "moviediary/internal/model/dto"
	service "moviediary/internal/service"
	utils "moviediary/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	service *service.TokenService
}

func NewTokenHandler(service *service.TokenService) *TokenHandler {
	return &TokenHandler{service: service}
}

func (tokenHandler *TokenHandler) RefreshToken(c *gin.Context) {
	var req model_dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := tokenHandler.service.RefreshToken(c.Request.Context(), req.Token)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "token", token, "Token refreshed successfully")

}
