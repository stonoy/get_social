package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stonoy/get_social/internal"
)

type MyCustomClaims struct {
	role string
	jwt.RegisteredClaims
}

func GenerateToken(secret string, user internal.User) (string, error) {
	claims := MyCustomClaims{
		string(user.Role),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "get_social",
			ID:        fmt.Sprintf("%v", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

func DecodeToken(secret, tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims.ID, nil
	} else {
		return "", errors.New("unknown claims type, cannot proceed")
	}
}
