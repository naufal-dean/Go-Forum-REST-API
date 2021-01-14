package posts

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response/data"
	"gorm.io/gorm"
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
// @Success  201  object  orm.Post  "Created Post"
// @Failure  400  object  data.ValidationErrorResponse  "Invalid Input Field(s) Error"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Failure  403  object  response.ErrorResponse  "Forbidden Error (Referenced Thread Not Exists)"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
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
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
		}

		// Check the referenced thread
		if err := a.DB.Where("id = ?", input.ThreadID).First(&orm.Thread{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusForbidden, "Thread referenced by 'thread_id' is not exists")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Create record
		post := orm.Post{Title: input.Title, Content: input.Content, UserID: claims.UserID, ThreadID: input.ThreadID}
		if err := a.DB.Create(&post).Error; err != nil {
			response.Error(w, http.StatusInternalServerError, "Create post failed")
			return
		}

		// Succeed
		response.JSON(w, http.StatusCreated, post)
	})
}
