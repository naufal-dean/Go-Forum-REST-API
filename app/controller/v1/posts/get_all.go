package posts

import (
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
)

func GetAll(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var posts []orm.Post
		a.DB.Find(&posts)

		// Succeed
		response.JSON(w, http.StatusOK, posts)
	})
}
