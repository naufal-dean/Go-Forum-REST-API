package posts

import (
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/data"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

func Delete(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	// Get params
		vars := mux.Vars(r)

		// Get record
		var post orm.Post
		if err := a.DB.Where("id = ?", vars["id"]).First(&post).Error; err != nil {
			response.JSON(w, http.StatusNotFound, data.ResourceNotFound())
			return
		}

		// Delete record
		a.DB.Delete(&post)

		// Succeed
		response.JSON(w, http.StatusNoContent, nil)
	})
}
