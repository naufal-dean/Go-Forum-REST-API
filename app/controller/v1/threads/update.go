package threads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
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
// @Failure  422  object  response.ErrorResponse  "Invalid Input Error JSON"
// @Failure  500  object  response.ErrorResponse  "Internal Server Error JSON"
// @Resource threads
// @Route /api/v1/threads/{id} [put]
func Update(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params and input
    	vars := mux.Vars(r)
		var input UpdateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

		// Get record
		var thread orm.Thread
		if err := a.DB.Where("id = ?", vars["id"]).First(&thread).Error; err != nil {
			response.Error(w, http.StatusNotFound, "Thread not found")
			return
		}

		// Check ownership
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
		}
		if claims.UserID != thread.UserID {
			response.Error(w, http.StatusForbidden, "You are not the owner of the thread")
			return
		}

    	// Update record
		a.DB.Model(&thread).Updates(orm.Thread{Name: input.Name})

    	// Succeed
		response.JSON(w, http.StatusOK, thread)
	})
}
