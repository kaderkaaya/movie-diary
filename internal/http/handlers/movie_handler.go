package handlers

import (
	model_dto "moviediary/internal/model/dto"
	service "moviediary/internal/service"
	utils "moviediary/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (movieHandler *MovieHandler) ListMovies(c *gin.Context) {
	var req model_dto.ListMoviesRequest
	
	if err := c.ShouldBindUri(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	movies, err := movieHandler.service.ListMovies(c.Request.Context(), req.MovieType, req.GenreID, req.Year, req.Page, req.Rating)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "movies", movies, "Movies discovered successfully")
}
