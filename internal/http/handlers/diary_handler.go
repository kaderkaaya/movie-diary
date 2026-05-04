package handlers

import (
	model_dto "moviediary/internal/model/dto"
	service "moviediary/internal/service"
	utils "moviediary/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiaryHandler struct {
	service *service.DiaryService
}

func NewDiaryHandler(service *service.DiaryService) *DiaryHandler {
	return &DiaryHandler{service: service}
}

func (diaryHandler *DiaryHandler) AddDiary(c *gin.Context) {
	var req model_dto.AddDiaryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	userIDValue, exists := c.Get("userID")
	if !exists {
		utils.Fail(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		utils.Fail(c, http.StatusUnauthorized, "Invalid user id")
		return
	}
	diary, err := diaryHandler.service.AddDiary(c.Request.Context(), userID, req.MovieId, req.Comment, req.Rating, req.WatchedAt.Time)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "diary", diary, "Diary added successfully")
}

func (diaryHandler *DiaryHandler) RemoveDiary(c *gin.Context) {
	var req model_dto.RemoveDiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	userIDValue, exists := c.Get("userID")
	if !exists {
		utils.Fail(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		utils.Fail(c, http.StatusUnauthorized, "Invalid user id")
		return
	}
	err := diaryHandler.service.RemoveDiary(c.Request.Context(), userID, req.MovieId)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "diary", nil, "Diary removed successfully")
}
