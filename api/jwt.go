package api

import (
	"context"
	"errors"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mvpoyatt/go-api/utils/logger"
)

var key = []byte("my_secret_key")

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func NewToken(email string) (string, error) {
	expiryTime := time.Now().Add(5 * time.Minute)
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

func ValidateToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		logger.Log.Infow("First thing broke")
	}

	if newClaims, ok := token.Claims.(*Claims); ok && token.Valid {
		return newClaims.Email, nil
	} else {
		return "", errors.New("Invalid login")
	}
}

func AuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			email, err := ValidateToken(req.Header().Get("jwtToken"))
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			} else {
				logger.Log.Info(email)
			}

			// if req.Header().Get(tokenHeader) == "" {
			// 	// Check token in handlers.
			// 	return nil, connect.NewError(
			// 		connect.CodeUnauthenticated,
			// 		errors.New("no token provided"),
			// 	)
			// }
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
