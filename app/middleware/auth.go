package middleware

import (
	"context"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

func Auth(a *core.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get claims from bearer token
			claims, err := auth.GetClaims(r)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, err.Error())
			}

			// Check tokens table
			err = a.DB.Where("user_id = ? AND token_uuid = ?", claims.UserID, claims.TokenUUID).First(&orm.Token{}).Error
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "Invalid token")
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
