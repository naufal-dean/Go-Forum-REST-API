package middleware

import (
	"context"
	"github.com/pkg/errors"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/auth"
	"github.com/naufal-dean/go-forum-rest-api/app/model/cerror"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"gorm.io/gorm"
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
				if errors.Is(err, gorm.ErrRecordNotFound) {
					response.Error(w, http.StatusUnauthorized, "Invalid token")
					return
				} else {
					panic(&cerror.DatabaseError{DBErr: err})
				}
			}

			// Save claims object as context
			ctx := context.WithValue(r.Context(), "claims", claims)
			r = r.WithContext(ctx)

			// Authenticated
			next.ServeHTTP(w, r)
		})
	}
}
