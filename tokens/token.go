package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	jwtScret           = []byte("sicommunitybackend")
	accessTokenExpiry  = time.Minute * 15
	refreshTokenExpiry = time.Hour * 24 * 7
)

type Token interface {
	GenerateToken(userID int32) (Token, error)
}

type AccessToken struct {
	gorm.Model
	ID             int32     `gorm:"primaryKey;" json:"id"`
	RegisterNumber int32     `gorm:"column:register_number" json:"register_number"`
	Token          string    `gorm:"column:token" json:"token"`
	ExpirationTime time.Time `gorm:"column:expiration_time" json:"expiration_time"`
}

func (AccessToken) TableName() string {
	return "access_tokens"
}

type RefreshToken struct {
	gorm.Model
	ID             int32     `gorm:"primaryKey;" json:"id"`
	RegisterNumber int32     `gorm:"column:register_number" json:"register_number"`
	Token          string    `gorm:"column:token" json:"token"`
	ExpirationTime time.Time `gorm:"column:expiration_time" json:"expiration_time"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (a AccessToken) GenerateToken(userID int32) (Token, error) {
	expirationTime := time.Now().Add(accessTokenExpiry)
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	accessJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessJwtToken.SignedString(jwtScret)
	if err != nil {
		return AccessToken{}, err
	}

	accessToken := AccessToken{
		RegisterNumber: userID,
		Token:          accessTokenString,
		ExpirationTime: expirationTime,
	}

	return accessToken, nil
}

func (r RefreshToken) GenerateToken(userID int32) (Token, error) {
	expirationTime := time.Now().Add(refreshTokenExpiry)
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}
	refreshJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshJwtToken.SignedString(jwtScret)
	if err != nil {
		return RefreshToken{}, err
	}

	refreshToken := RefreshToken{
		RegisterNumber: userID,
		Token:          refreshTokenString,
		ExpirationTime: expirationTime,
	}

	return refreshToken, nil
}
