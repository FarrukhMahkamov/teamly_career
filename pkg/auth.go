package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AuthClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	TokenTTL = 14 * 24 * time.Hour
)

func InitEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

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

	signingKey := os.Getenv("SIGNINGKEY")

	SignedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		logrus.Errorf("failed to sign JWT token: %s", err)
		return "", err
	}

	return SignedToken, nil
}

func ParseToken(token string) (int64, error) {
	claims := &AuthClaims{}

	// replace with your own secret key
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*AuthClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserID, nil
}
