package posts

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
	Title   string `json:"title"`
	Content string `json:"content"`
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
		var post orm.Post
		if err := a.DB.Where("id = ?", vars["id"]).First(&post).Error; err != nil {
			response.JSON(w, http.StatusNotFound, data.ResourceNotFound())
			return
		}

    	// Update record
		a.DB.Model(&post).Updates(orm.Post{Title: input.Title, Content: input.Content})

    	// Succeed
		response.JSON(w, http.StatusOK, post)
	})
}
