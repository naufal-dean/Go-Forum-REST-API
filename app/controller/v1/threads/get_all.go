package threads

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

func GetAll(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var threads []orm.Thread
		a.DB.Find(&threads)

		// Succeed
		response.JSON(w, http.StatusOK, threads)
	})
}
