package threads

import (
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
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
