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
