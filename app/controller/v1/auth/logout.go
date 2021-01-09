package auth

import (
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/lib/auth"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"github.com/pkg/errors"
	"net/http"
)

func Logout(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
		}

		// Delete token from table
		ormToken := &orm.Token{}
		err := a.DB.Where("user_id = ? AND token_uuid = ?", claims.UserID, claims.TokenUUID).First(&ormToken).Error
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "Invalid token value last")
			return
		}
		a.DB.Delete(&ormToken)

		// Succeed
		// TODO: create response object
		response.JSON(w, http.StatusOK, "Logout succeed")
	})
}
