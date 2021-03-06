package threads

import (
	"encoding/json"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/auth"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"github.com/naufal-dean/go-forum-rest-api/app/response/data"
	"net/http"
)

type CreateInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// @Title Create a thread.
// @Description Create a new thread.
// @Param  thread  body  CreateInput  true  "Thread data."
// @Success  201  object  orm.Thread  "Created Thread"
// @Failure  400  object  data.ValidationErrorResponse  "Invalid Input Field(s) Error"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
// @Resource threads
// @Route /api/v1/threads [post]
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
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
		}

		// Create record
		thread := orm.Thread{Name: input.Name, Description: input.Description, UserID: claims.UserID}
		if err := a.DB.Create(&thread).Error; err != nil {
			response.Error(w, http.StatusInternalServerError, "Create thread failed")
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, thread)
	})
}
