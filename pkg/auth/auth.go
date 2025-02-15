package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"shopline/pkg/constants"
	"time"
)

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash verifies a plain-text password against a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for the given user ID and admin status
func GenerateJWT(userID uint, isAdmin bool) (string, error) {
	secretKey := constants.JWTSecretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(constants.JWTDuration).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
