package posts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gorm.io/gorm"
	"net/http"
)

type UpdateInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// @Title Update a post.
// @Description Update a post with ID.
// @Param  id  path  int  true  "Post ID."
// @Param  post  body  UpdateInput  optional  "Post new data."
// @Success  200  object  orm.Post  "Updated Post Data"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Failure  403  object  response.ErrorResponse  "Forbidden Error (Non Owner)"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
// @Resource posts
// @Route /api/v1/posts/{id} [put]
func Update(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get input
		var input UpdateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

		// Get record
		var post orm.Post
		if err := a.DB.Where("id = ?", id).First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusNotFound, "Post not found")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Check ownership
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
		}
		if claims.UserID != post.UserID {
			response.Error(w, http.StatusForbidden, "You are not the owner of the post")
			return
		}

		// Update record
		a.DB.Model(&post).Updates(orm.Post{Title: input.Title, Content: input.Content})

		// Succeed
		response.JSON(w, http.StatusOK, post)
	})
}
