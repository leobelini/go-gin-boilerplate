package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims customizadas
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Gera token JWT
func GenerateJWT(userID, jwtKey string) (string, error) {
	fmt.Println("Generating JWT for user ID:", userID) // Log do userID
	fmt.Println("Using JWT Key:", jwtKey)              // Log do jwtKey
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // expira em 24h
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

// Valida token JWT e retorna userID
func ValidateJWT(tokenString, jwtKey string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.UserID, nil
}
