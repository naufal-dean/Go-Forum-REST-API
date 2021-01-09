package posts

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
	Title    string `json:"title"`
	Content  string `json:"content"`
	ThreadID uint   `json:"thread_id"`
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
		post := orm.Post{Title: input.Title, Content: input.Content, UserID: claims.UserID, ThreadID: input.ThreadID}
		if err := a.DB.Create(&post).Error;
			err != nil {
			response.JSON(w, http.StatusInternalServerError, data.CustomError("Create post failed"))
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, post)
	})
}
