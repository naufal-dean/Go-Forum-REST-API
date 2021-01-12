package posts

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get a post.
// @Description Get a post with ID.
// @Param  id  path  int  true  "Post ID."
// @Success  200  object  orm.Post  "Post JSON"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error JSON"
// @Resource posts
// @Route /api/v1/posts/{id} [get]
func GetOne(a *core.App) http.Handler {
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
			response.Error(w, http.StatusNotFound, "Post not found")
			return
		}

		// Succeed
		response.JSON(w, http.StatusOK, post)
	})
}
