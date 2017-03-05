package utils

import (
	"fmt"
	"time"

	"../config"
	"../models"
	"github.com/dgrijalva/jwt-go"
)

// Claims ...
type Claims struct {
	User models.User `json:"user"`
	// recommended having
	jwt.StandardClaims
}

// CreateUserToken ...
func CreateUserToken(user models.User, cfg config.Config) (token string, expiresat time.Time, err error) {
	expiresAt := time.Now().Add(time.Hour * 10)

	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "localhost:8080",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := jwtToken.SignedString([]byte(cfg.JWTSecret))

	return tokenString, expiresAt, err
}

// DecodeUserFromJWT ...
func DecodeUserFromJWT(tokenString string, cfg config.Config) (user *models.User, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}

		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &claims.User, nil
	}

	return nil, err
}
