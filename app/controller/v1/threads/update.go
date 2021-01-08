package threads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
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
