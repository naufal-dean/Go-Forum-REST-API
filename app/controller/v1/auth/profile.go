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

// @Title Profile.
// @Description Get the profile of authenticated user.
// @Success  200  object  orm.User  "Authenticated User Data"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Resource auth
// @Route /api/v1/profile [post]
func Profile(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
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
