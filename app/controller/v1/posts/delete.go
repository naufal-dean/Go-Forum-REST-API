package posts

import (
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

// TODO: add no content return object

// @Title Delete a post.
// @Description Delete a post with ID.
// @Param  id  path  int  true  "Post ID."
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error JSON"
// @Resource posts
// @Route /api/v1/posts/{id} [delete]
func Delete(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

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

		// Delete record
		a.DB.Delete(&post)

		// Succeed
		response.JSON(w, http.StatusNoContent, nil)
	})
}
