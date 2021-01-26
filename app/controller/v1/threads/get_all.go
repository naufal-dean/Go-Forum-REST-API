package threads

import (
	"fmt"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"net/http"
)

// @Title Get all threads.
// @Description Get all thread in database.
// @Param  search  query  string  optional  "Search query for thread name."
// @Success  200  array  []orm.Thread  "Array of Thread JSON"
// @Resource threads
// @Route /api/v1/threads [get]
func GetAll(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get records
		var threads []orm.Thread
		search := r.URL.Query().Get("search")
		if search == "" {
			a.DB.Find(&threads)
		} else {
			a.DB.Where("name LIKE ?", fmt.Sprintf("%%%s%%", search)).Find(&threads)
		}

		// Succeed
		response.JSON(w, http.StatusOK, threads)
	})
}
