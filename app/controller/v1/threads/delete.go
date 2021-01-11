package threads

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/data"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// TODO: add no content return object

// @Title Delete a thread.
// @Description Delete a thread with ID.
// @Param  id  path  int  true  "Thread ID."
// @Failure  404  object  data.ErrorResponse  "Resource Not Found Error JSON"
// @Resource threads
// @Route /api/v1/threads/{id} [delete]
func Delete(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	// Get params
		vars := mux.Vars(r)

		// Get record
		var thread orm.Thread
		if err := a.DB.Where("id = ?", vars["id"]).First(&thread).Error; err != nil {
			response.JSON(w, http.StatusNotFound, data.ResourceNotFound())
			return
		}

		// Delete record
		a.DB.Delete(&thread)

		// Succeed
		response.JSON(w, http.StatusNoContent, nil)
	})
}
