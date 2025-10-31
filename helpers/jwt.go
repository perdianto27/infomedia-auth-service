package helpers

import (
	"os"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email 	string `json:"email"`
	Name   		string`json:"name"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Generate Access Token (Expired 15 Menit)
func GenerateTokenJWT(email string, name string) (string, error) {
	claims := JwtCustomClaims{
		Email: 	email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET_JWT")

	return token.SignedString([]byte(secret))
}

// Generate Refresh Token (Expired 7 Hari)
func GenerateRefreshToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": 	email,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET_REFRESH_JWT")

	return token.SignedString([]byte(secret))
}

func VerifyAccessToken(tokenString string) (*JwtCustomClaims, error) {
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		return nil, errors.New("SECRET_JWT is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}