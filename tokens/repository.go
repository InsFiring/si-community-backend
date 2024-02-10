package tokens

import (
	"errors"
	"fmt"
	user "si-community/users"
	"time"

	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (t *TokenRepository) GetOrCreateAccessToken(userResponseDto *user.UserResponseDto) (*user.UserResponseDto, error) {
	var accessToken AccessToken
	now := time.Now()

	result := t.db.Table("access_tokens").
		Where(&AccessToken{RegisterNumber: userResponseDto.RegisterNumber}).
		Last(&accessToken)
	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Access token 조회 에러")
		return userResponseDto, result.Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || accessToken.ExpirationTime.Before(now) {
		newToken, err := accessToken.GenerateToken(userResponseDto.RegisterNumber)
		if err != nil {
			return userResponseDto, err
		}

		newAccessToken := newToken.(AccessToken)
		if err := t.db.Create(&newAccessToken).Error; err != nil {
			return userResponseDto, err
		}

		userResponseDto.AccessToken = newAccessToken.Token

		return userResponseDto, nil
	}

	userResponseDto.AccessToken = accessToken.Token

	return userResponseDto, nil
}

func (t *TokenRepository) GetOrCreateRefreshToken(userResponseDto *user.UserResponseDto) (*user.UserResponseDto, error) {
	var refreshToken RefreshToken
	now := time.Now()

	result := t.db.Table("refresh_tokens").
		Where(&RefreshToken{RegisterNumber: userResponseDto.RegisterNumber}).
		Last(&refreshToken)
	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Refresh token 조회 에러")
		return userResponseDto, result.Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || refreshToken.ExpirationTime.Before(now) {
		newToken, err := refreshToken.GenerateToken(userResponseDto.RegisterNumber)
		if err != nil {
			return userResponseDto, err
		}

		newRefreshToken := newToken.(RefreshToken)
		if err := t.db.Create(&newRefreshToken).Error; err != nil {
			return userResponseDto, err
		}

		userResponseDto.RefreshToken = newRefreshToken.Token

		return userResponseDto, nil
	}

	userResponseDto.RefreshToken = refreshToken.Token

	return userResponseDto, nil
}
