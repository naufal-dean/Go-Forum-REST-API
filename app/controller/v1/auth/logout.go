package auth

import (
	"github.com/pkg/errors"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/auth"
	"github.com/naufal-dean/go-forum-rest-api/app/model/cerror"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"gorm.io/gorm"
	"net/http"
)

// @Title Logout.
// @Description Invalidate current token.
// @Success  200  object  response.SuccessResponse  "Logout Succeed"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Resource auth
// @Route /api/v1/logout [post]
func Logout(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
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
