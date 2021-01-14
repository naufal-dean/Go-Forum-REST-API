package users

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gorm.io/gorm"
	"net/http"
)

// @Title Get user's thread.
// @Description Get all threads owned by user.
// @Param  id  path  int  true  "User ID."
// @Success  200  array  orm.Thread  "Array of User's Thread JSON"
// @Resource users/threads
// @Route /api/v1/users/{id}/threads [get]
func GetThreads(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get records
		var user orm.User
		if err := a.DB.Preload("Threads").Where("id = ?", id).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusNotFound, "User not found")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Succeed
		response.JSON(w, http.StatusOK, user.Threads)
	})
}
