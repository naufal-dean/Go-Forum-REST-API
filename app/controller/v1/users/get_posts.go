package users

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get user's post.
// @Description Get all posts owned by user.
// @Param  id  path  int  true  "User ID."
// @Success  200  array  orm.Post  "Array of User's Post JSON"
// @Resource users/posts
// @Route /api/v1/users/{id}/posts [get]
func GetPosts(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get records
		var posts []orm.Post
		a.DB.Where("user_id = ?", id).Find(&posts)

		// Succeed
		response.JSON(w, http.StatusOK, posts)
	})
}
