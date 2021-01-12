package threads

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/util"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get thread's post.
// @Description Get all posts in the thread.
// @Param  id  path  int  true  "Thread ID."
// @Success  200  array  orm.Post  "Array of Thread's Post JSON"
// @Resource threads/posts
// @Route /api/v1/threads/{id}/posts [get]
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
		a.DB.Where("thread_id = ?", id).Find(&posts)

		// Succeed
		response.JSON(w, http.StatusOK, posts)
    })
}

