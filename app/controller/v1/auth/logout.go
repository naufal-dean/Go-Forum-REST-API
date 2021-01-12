package auth

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

// TODO: create succeed response

// @Title Logout.
// @Description Invalidate current token.
// @Success  200  object  response.SuccessResponse  "Logout Succeed JSON"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error JSON"
// @Resource auth
// @Route /api/v1/logout [post]
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusUnauthorized, "Invalid token")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}
		a.DB.Delete(&ormToken)

		// Succeed
		response.Success(w, http.StatusOK, "Logout succeed")
	})
}
