package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTni yaratish uchun kalit
var JWT_SECRET_KEY = []byte("secret")

// Token yaratish
func GenerateJWT(userID int) (string, error) {
	// Tokenni yaratish
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Tokenni imzolash
	return token.SignedString(JWT_SECRET_KEY)
}
