package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func init() {
	if secretKey == "" {
		secretKey = "secret_key_for_testing_purpose"
	}
}

type Claims struct {
	UserID  string `json:"user_id"`
	ModelID string `json:"model_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, modelID string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime)
	claims := &Claims{
		UserID:  userID,
		ModelID: modelID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func Authenticate(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
