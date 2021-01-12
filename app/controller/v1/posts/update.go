package posts

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
	Title   string `json:"title"`
	Content string `json:"content"`
}

// @Title Update a post.
// @Description Update a post with ID.
// @Param  id  path  int  true  "Post ID."
// @Param  post  body  UpdateInput  optional  "Post new data."
// @Success  200  object  orm.Post  "Updated Post JSON"
// @Failure  422  object  response.ErrorResponse  "Invalid Input Error JSON"
// @Failure  500  object  response.ErrorResponse  "Internal Server Error JSON"
// @Resource posts
// @Route /api/v1/posts/{id} [put]
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
		var post orm.Post
		if err := a.DB.Where("id = ?", vars["id"]).First(&post).Error; err != nil {
			response.Error(w, http.StatusNotFound, "Post not found")
			return
		}

		// Check ownership
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			panic(errors.New("invalid claims context"))
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
