package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtScret           = []byte("sicommunitybackend")
	accessTokenExpiry  = time.Minute * 15
	refreshTokenExpiry = time.Hour * 24 * 7
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func GenerateTokens(userID int32) (*Token, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(accessTokenExpiry).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtScret)
	if err != nil {
		return &Token{}, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(refreshTokenExpiry).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtScret)
	if err != nil {
		return &Token{}, err
	}

	token := &Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return token, nil
}
