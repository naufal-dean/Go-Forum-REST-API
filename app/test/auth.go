package test

import (
	"context"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/auth"
	"net/http"
)

func ActAsUser(r *http.Request, userID int) *http.Request {
	// Inject claims context
	ctx := context.WithValue(r.Context(), "claims", &auth.TokenClaims{
		UserID:    TokensData[userID-1].UserID,
		TokenUUID: TokensData[userID-1].TokenUUID,
	})
	return r.WithContext(ctx)
}
