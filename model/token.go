package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Key string `json:"key"`
	jwt.RegisteredClaims
}
