package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/lib/auth"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
	"os"
	"strings"
)

func Auth(a *core.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token
			authorizationHeader := r.Header.Get("Authorization")
			arr := strings.Split(authorizationHeader, " ")
			tokenString := ""
			if len(arr) == 2 && arr[0] == "Bearer" {
				tokenString = arr[1]
			} else {
				// TODO: create response object
				response.JSON(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Parse token
			token, err := jwt.ParseWithClaims(tokenString, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Check token signing method
				method, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok || method != jwt.SigningMethodHS256 {
					// TODO: create response object
					return nil, fmt.Errorf("invalid token signing method")
				}
				// Return
				return []byte(os.Getenv("APP_SECRET")), nil
			})
			if err != nil {
				// TODO: create response object
				response.JSON(w, http.StatusUnauthorized, err.Error())
				return
			}

			// Check claims
			claims, ok := token.Claims.(*auth.TokenClaims)
			if !ok || !token.Valid {
				// TODO: create response object
				response.JSON(w, http.StatusUnauthorized, "Invalid token claims")
				return
			}

			// Check tokens table
			err = a.DB.Where("user_id = ? AND token_uuid = ?", claims.UserID, claims.TokenUUID).First(&orm.Token{}).Error
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, "Invalid token value last")
				return
			}

			// Save claims object as context
			ctx := context.WithValue(r.Context(), "claims", claims)
			r = r.WithContext(ctx)

			// Authenticated
			next.ServeHTTP(w, r)
		})
	}
}
