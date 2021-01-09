package posts

import (
	"encoding/json"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
)

type CreateInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"` // TODO: remove
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

		// Create record
		// TODO: get user id from auth
		post := orm.Post{Title: input.Title, Content: input.Content, UserID: input.UserID, ThreadID: input.ThreadID}
		if err := a.DB.Create(&post).Error;
			err != nil {
			response.JSON(w, http.StatusInternalServerError, data.CustomError("Create post failed"))
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, post)
	})
}
