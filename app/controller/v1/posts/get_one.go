package posts

import (
	"github.com/gorilla/mux"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
)

func GetOne(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		vars := mux.Vars(r)

		// Get record
		var post orm.Post
		if err := a.DB.Where("id = ?", vars["id"]).First(&post).Error; err != nil {
			response.JSON(w, http.StatusNotFound, data.ResourceNotFound())
			return
		}

		// Succeed
		response.JSON(w, http.StatusOK, post)
	})
}
