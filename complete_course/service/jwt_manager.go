package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


type JWTManager struct {
	secretKey string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role string `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

