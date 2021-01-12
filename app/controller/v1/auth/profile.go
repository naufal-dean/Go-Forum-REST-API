package auth

import (
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gorm.io/gorm"
	"net/http"
)

// @Title Profile.
// @Description Get the profile of authenticated user.
// @Success  200  object  orm.User  "Authenticated User JSON"
// @Resource auth
// @Route /api/v1/profile [post]
func Profile(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
		}

    	// Get record
		var user orm.User
		if err := a.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(errors.New("No user record related to a valid token claim"))
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Succeed
		response.JSON(w, http.StatusOK, user)
    })
}

