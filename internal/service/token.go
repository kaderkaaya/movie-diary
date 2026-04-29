package service

import (
	"context"
	config "moviediary/internal/config"
	model "moviediary/internal/model"
	repository "moviediary/internal/repository"
	apperror "moviediary/pkg/apperror"
	utils "moviediary/pkg/utils"
	"time"
)

type TokenService struct {
	tokenRepository *repository.TokenRepository
}

func NewTokenService(tokenRepository *repository.TokenRepository) *TokenService {
	return &TokenService{tokenRepository: tokenRepository}
}

func (tokenService *TokenService) RefreshToken(ctx context.Context, token string) (*model.UserTokens, error) {
	if token == "" {
		return nil, apperror.ErrTokenEmpty
	}

	userToken, err := tokenService.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if userToken.ExpiresAt.Before(time.Now()) {
		return nil, apperror.ErrTokenExpired
	}

	newToken, err := utils.GenerateJWT(userToken.UserID, config.Load().JwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	userToken, err = tokenService.tokenRepository.UpdateUserToken(ctx, userToken.UserID, newToken)
	if err != nil {
		return nil, err
	}
	return userToken, nil
}
