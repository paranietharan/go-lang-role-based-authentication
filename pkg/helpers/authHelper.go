package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("s3cReT!@9xYz8&$%PqLmN*UvW0aBcD%")

type Claims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"usertype"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, email, userType string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Email:    email,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day expiration
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
