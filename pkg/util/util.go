package util

import (
	"errors"
	"os"
	"reflect"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func IsZero(x any) bool {
	if x == nil {
		return true
	}
	return reflect.ValueOf(x).IsZero()
}

type JWTClaims struct {
	UserID  int64  `json:"uid"`
	Account string `json:"acc"`
	jwt.RegisteredClaims
}

func jwtSecret() []byte {
	s := os.Getenv("ADMIN_JWT_SECRET")
	if s == "" {
		s = "admin-service-secret"
	}
	return []byte(s)
}

// GenerateJWT creates a signed JWT token for a user
func GenerateJWT(userID int64, account string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:  userID,
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "admin-service",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret())
}

// ParseJWT verifies token and returns claims
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
