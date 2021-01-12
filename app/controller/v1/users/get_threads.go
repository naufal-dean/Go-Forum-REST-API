package users

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
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
		var threads []orm.Thread
		a.DB.Where("user_id = ?", id).Find(&threads)

		// Succeed
		response.JSON(w, http.StatusOK, threads)
	})
}
