package threads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/data"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

type UpdateInput struct {
	Name   string `json:"name"`
}

// @Title Update a thread.
// @Description Update a thread with ID.
// @Param  id  path  int  true  "Thread ID."
// @Param  thread  body  UpdateInput  optional  "Thread new data."
// @Success  200  object  orm.Thread  "Updated Thread JSON"
// @Failure  422  object  data.ErrorResponse  "Invalid Input Error JSON"
// @Failure  500  object  data.ErrorResponse  "Internal Server Error JSON"
// @Resource threads
// @Route /api/v1/threads/{id} [put]
func Update(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params and input
    	vars := mux.Vars(r)
		var input UpdateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, data.CustomError("Invalid input"))
			return
		}
		defer r.Body.Close()

		// Get record
		var thread orm.Thread
		if err := a.DB.Where("id = ?", vars["id"]).First(&thread).Error; err != nil {
			response.JSON(w, http.StatusNotFound, data.ResourceNotFound())
			return
		}

    	// Update record
		a.DB.Model(&thread).Updates(orm.Thread{Name: input.Name})

    	// Succeed
		response.JSON(w, http.StatusOK, thread)
	})
}
