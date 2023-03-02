package api

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
)

var key = []byte("change_this")

// Set how long login is valid for
var tokenDuration = 5 * time.Minute

// How close to end of session to create new token
var refreshWindow = tokenDuration / 5

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func NewToken(email string) (string, error) {
	expiryTime := time.Now().Add(tokenDuration)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ValidateToken(tokenString string) (string, string, error) {
	if tokenString == "" {
		return "", "", errors.New("authentication token not found")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", "", errors.New("problem parsing authentication token")
	}

	if newClaims, ok := token.Claims.(*Claims); ok && token.Valid {
		refreshCutoff := newClaims.ExpiresAt.Time.Add(-refreshWindow)
		if time.Now().After(refreshCutoff) {
			// Token about to expire, send back a new token
			if refreshToken, err := NewToken(newClaims.Email); err != nil {
				// New token didn't work but still validate this request
				return newClaims.Email, "", nil
			} else {
				return newClaims.Email, refreshToken, nil
			}
		} else {
			// Token not about to expire, don't send back a new one
			return newClaims.Email, "", nil
		}
	} else {
		return "", "", errors.New("invalid login")
	}
}

func AuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			path := strings.Split(req.Spec().Procedure, "/")
			endpoint := path[len(path)-1]
			if endpoint == "PutUser" || endpoint == "LoginUser" {
				return next(ctx, req)
			}
			token := req.Header().Get("Authorization")
			email, refreshToken, err := ValidateToken(token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			} else {
				ctx = context.WithValue(ctx, "email", email)
				if refreshToken != "" {
					ctx = context.WithValue(ctx, "refreshToken", refreshToken)
				}
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
