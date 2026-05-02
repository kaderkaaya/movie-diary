package service

import (
	"context"
	"time"

	config "moviediary/internal/config"
	model "moviediary/internal/model"
	repository "moviediary/internal/repository"
	apperror "moviediary/pkg/apperror"
	utils "moviediary/pkg/utils"
)

type TokenService struct {
	tokenRepository *repository.TokenRepository
	userRepository  *repository.UserRepository
}

func NewTokenService(tokenRepository *repository.TokenRepository, userRepository *repository.UserRepository) *TokenService {
	return &TokenService{tokenRepository: tokenRepository, userRepository: userRepository}
}

func (tokenService *TokenService) RefreshToken(ctx context.Context, token string) (*model.UserTokens, error) {
	if token == "" {
		return nil, apperror.ErrTokenEmpty
	}

	userToken, err := tokenService.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	//user, err := tokenService.userRepository.FindByID(ctx, userToken.UserID)
	//if err != nil {
	//	return nil, err
	//}
	if err != nil {
		if err != apperror.ErrTokenNotFound {
			return nil, err
		}
		claims, errJWT := utils.ParseJWT(token, config.Load().JwtSecret)
		if errJWT != nil {
			return nil, apperror.ErrTokenNotFound
		}
		expAt := time.Now().Add(7 * 24 * time.Hour)
		if claims.ExpiresAt != nil {
			expAt = claims.ExpiresAt.Time
		}
		userToken = &model.UserTokens{
			UserID:    claims.UserID,
			Token:     token,
			ExpiresAt: expAt,
		}
	}

	if userToken.ExpiresAt.Before(time.Now()) {
		return nil, apperror.ErrTokenExpired
	}

	newToken, err := utils.GenerateJWT(userToken.UserID, config.Load().JwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	return tokenService.tokenRepository.UpdateUserToken(ctx, userToken.UserID, newToken)
}
