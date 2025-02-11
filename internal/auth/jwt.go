package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"shopline/pkg/constants"
	"time"
)

var jwtKey = []byte("secret") // TODO get the secret from .env file

// Claims represents the JWT claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for the given user ID.
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(constants.JWTDuration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWTSecretKey))
}

// ValidateToken validates a JWT token and returns the claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
