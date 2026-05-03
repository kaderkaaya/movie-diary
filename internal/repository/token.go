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
	if err := tokenRepository.db.WithContext(ctx).Preload("User").Where("token = ?", token).First(&userToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrTokenNotFound
		}
		return nil, err
	}
	return &userToken, nil
}

func (tokenRepository *TokenRepository) UpdateUserToken(ctx context.Context, userID uint, token string) (*model.UserTokens, error) {
	exp := time.Now().Add(7 * 24 * time.Hour)
	var ut model.UserTokens
	err := tokenRepository.db.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").First(&ut).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ut = model.UserTokens{
			UserID:    userID,
			Token:     token,
			ExpiresAt: exp,
		}
		if err := tokenRepository.db.WithContext(ctx).Create(&ut).Error; err != nil {
			return nil, err
		}
		return tokenRepository.userTokenByID(ctx, ut.ID)
	}
	if err != nil {
		return nil, err
	}
	ut.Token = token
	ut.ExpiresAt = exp
	if err := tokenRepository.db.WithContext(ctx).Save(&ut).Error; err != nil {
		return nil, err
	}
	return tokenRepository.userTokenByID(ctx, ut.ID)
}

func (tokenRepository *TokenRepository) userTokenByID(ctx context.Context, id uint) (*model.UserTokens, error) {
	var out model.UserTokens
	if err := tokenRepository.db.WithContext(ctx).Preload("User").First(&out, id).Error; err != nil {
		return nil, err
	}
	return &out, nil
}
