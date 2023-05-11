package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	salt       = "teamly"
	SigningKey = "teamly"
	TokenTTL   = 14 * 24 * time.Hour
)

// GenerateToken generates a new JWT token with the given user ID
func GenerateToken(userID int64) (string, error) {
	claims := &AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "teamly",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SigningKey))
}
