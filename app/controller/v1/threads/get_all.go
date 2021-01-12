package threads

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

// @Title Get all threads.
// @Description Get all thread in database.
// @Success  200  array  []orm.Thread  "Array of Thread JSON"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error JSON"
// @Resource threads
// @Route /api/v1/threads [get]
func GetAll(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var threads []orm.Thread
		a.DB.Find(&threads)

		// Succeed
		response.JSON(w, http.StatusOK, threads)
	})
}
