package threads

import (
	"encoding/json"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/lib/auth"
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"github.com/pkg/errors"
	"net/http"
)

type CreateInput struct {
	Name   string `json:"name"`
}

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
