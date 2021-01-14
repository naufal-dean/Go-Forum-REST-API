package posts

import (
	"fmt"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get all posts.
// @Description Get all post in database.
// @Param  search  query  string  optional  "Search query for post title."
// @Success  200  array  []orm.Post  "Array of Post"
// @Resource posts
// @Route /api/v1/posts [get]
func GetAll(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var posts []orm.Post
		search := r.URL.Query().Get("search")
		if search == "" {
			a.DB.Find(&posts)
		} else {
			a.DB.Where("title LIKE ?", fmt.Sprintf("%%%s%%", search)).Find(&posts)
		}

		// Succeed
		response.JSON(w, http.StatusOK, posts)
	})
}
