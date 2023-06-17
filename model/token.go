package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}
