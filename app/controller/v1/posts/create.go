package posts

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response/data"
	"net/http"
)

type CreateInput struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ThreadID uint   `json:"thread_id" validate:"required"`
}

// @Title Create a post.
// @Description Create a new post to a thread.
// @Param  post  body  CreateInput  true  "Post data."
// @Success  201  object  orm.Post  "Created Post JSON"
// @Failure  422  object  response.ErrorResponse  "Invalid Input Error JSON"
// @Failure  500  object  response.ErrorResponse  "Internal Server Error JSON"
// @Resource posts
// @Route /api/v1/posts [post]
func Create(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input CreateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

		// Validate input
		err := a.Validate.Struct(input)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, data.NewValidationErrorResponse(err))
			return
		}

		// Get claims context
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
		}

		// Create record
		post := orm.Post{Title: input.Title, Content: input.Content, UserID: claims.UserID, ThreadID: input.ThreadID}
		if err := a.DB.Create(&post).Error;
			err != nil {
			response.Error(w, http.StatusInternalServerError, "Create post failed")
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, post)
	})
}
