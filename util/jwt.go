package util

import (
	"ohsundosun-api/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetAccessToken(user *model.User) string {
	accessClaims := model.TokenClaims{
		Key: user.Key,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	accessAt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, _ := accessAt.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	return accessToken
}

func GetRefreshToken(user *model.User) string {
	refreshClaims := model.TokenClaims{
		Key: user.Key,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 14)),
		},
	}
	refreshAt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, _ := refreshAt.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))
	return refreshToken
}
