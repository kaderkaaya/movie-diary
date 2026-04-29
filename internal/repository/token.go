package repository

import (
	"context"
	"errors"
	model "moviediary/internal/model"
	"moviediary/pkg/apperror"
	"time"

	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (tokenRepository *TokenRepository) CreateUserToken(ctx context.Context, UserID uint, token string) (*model.UserTokens, error) {
	userToken := &model.UserTokens{
		UserID:    UserID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := tokenRepository.db.WithContext(ctx).Create(userToken).Error; err != nil {
		return nil, err
	}
	return &model.UserTokens{
		UserID: UserID,
		Token:  token,
	}, nil
}

func (tokenRepository *TokenRepository) FindByToken(ctx context.Context, token string) (*model.UserTokens, error) {
	var userToken model.UserTokens
	if err := tokenRepository.db.WithContext(ctx).Where("token = ?", token).First(&userToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrTokenNotFound
		}
		return nil, err
	}
	return &userToken, nil
}

func (tokenRepository *TokenRepository) UpdateUserToken(ctx context.Context, UserID uint, token string) (*model.UserTokens, error) {
	if err := tokenRepository.db.WithContext(ctx).Model(&model.UserTokens{}).Where("user_id = ?", UserID).Update("token", token).Error; err != nil {
		return nil, err
	}
	return &model.UserTokens{
		UserID: UserID,
		Token:  token,
	}, nil
}
