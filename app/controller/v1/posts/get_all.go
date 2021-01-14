package posts

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get all posts.
// @Description Get all post in database.
// @Success  200  array  []orm.Post  "Array of Post"
// @Resource posts
// @Route /api/v1/posts [get]
func GetAll(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var posts []orm.Post
		a.DB.Find(&posts)

		// Succeed
		response.JSON(w, http.StatusOK, posts)
	})
}
