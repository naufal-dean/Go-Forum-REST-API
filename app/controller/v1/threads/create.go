package threads

import (
	"encoding/json"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/data"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"github.com/pkg/errors"
	"net/http"
)

type CreateInput struct {
	Name   string `json:"name"`
}

// @Title Create a thread.
// @Description Create a new thread.
// @Param  thread  body  CreateInput  true  "Thread data."
// @Success  201  object  orm.Thread  "Created Thread JSON"
// @Failure  422  object  data.ErrorResponse  "Invalid Input Error JSON"
// @Failure  500  object  data.ErrorResponse  "Internal Server Error JSON"
// @Resource threads
// @Route /api/v1/threads [post]
func Create(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input CreateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, data.CustomError("Invalid input"))
			return
		}
		defer r.Body.Close()

		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
		}

		// Create record
		thread := orm.Thread{Name: input.Name, UserID: claims.UserID}
		if err := a.DB.Create(&thread).Error;
			err != nil {
			response.JSON(w, http.StatusInternalServerError, data.CustomError("Create thread failed"))
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, thread)
	})
}
