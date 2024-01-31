package middleware

import (
	"context"
	"modular-monolithic/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func MiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := validateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach the user information from the token to the request context for later use
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*jwt.StandardClaims).Subject)
		next(w, r.WithContext(ctx))
	}
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

func validateToken(tokenString string) (*jwt.Token, error) {
	// LOAD CONFIG DATA
	config := config.Get()

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppApiKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
