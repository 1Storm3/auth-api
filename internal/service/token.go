package service

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type EmailClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) GenerateToken(jwtKey []byte, userID, role string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
