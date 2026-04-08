package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shailendrapawar/book-store/internal/adapters"
)

// type Claims struct {
// 	UserID string `json:"user_id"`
// 	Role   string `json:"role"`
// 	jwt.RegisteredClaims
// }

func GenerateJwtToken(userID, role, secret string, expiryHours int) (string, error) {
	claims := &adapters.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiryHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (*adapters.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &adapters.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*adapters.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
