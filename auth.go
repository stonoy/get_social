package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("empty header -> %v", header)
	}

	headerArr := strings.Fields(header)

	if headerArr[0] == "Bearer" || len(headerArr) == 2 {
		return headerArr[1], nil
	} else {
		return "", errors.New("No valid token provided")
	}
}

type authCheckerType func(http.ResponseWriter, *http.Request, internal.User)

func (cfg *apiConfig) authChecker(mainFunc authCheckerType) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token
		token, err := GetTokenFromHeader(r)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in GetTokenFromHeader -> %v", err))
			return
		}

		// validate token
		userIdStr, err := DecodeToken(cfg.jwt_secret, token)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in DecodeToken -> %v", err))
			return
		}

		// parse user id
		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
			return
		}

		// get user from id
		user, err := cfg.db.GetUserById(r.Context(), userId)
		if err != nil {
			if err == sql.ErrNoRows {
				respWithError(w, 400, "No such user exist")
				return
			} else {
				respWithError(w, 500, fmt.Sprintf("error in GetUserById -> %v", err))
				return
			}
		}

		// call main func
		mainFunc(w, r, user)
	}
}
