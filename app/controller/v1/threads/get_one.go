package threads

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get a thread.
// @Description Get a thread with ID.
// @Param  id  path  int  true  "Thread ID."
// @Success  200  object  orm.Thread  "Thread JSON"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error JSON"
// @Resource threads
// @Route /api/v1/threads/{id} [get]
func GetOne(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get record
		var thread orm.Thread
		if err := a.DB.Where("id = ?", id).First(&thread).Error; err != nil {
			response.Error(w, http.StatusNotFound, "Thread not found")
			return
		}

		// Succeed
		response.JSON(w, http.StatusOK, thread)
	})
}
