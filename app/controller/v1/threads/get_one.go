package threads

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/util"
	"github.com/naufal-dean/go-forum-rest-api/app/model/cerror"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"gorm.io/gorm"
	"net/http"
)

// @Title Get a thread.
// @Description Get a thread with ID.
// @Param  id  path  int  true  "Thread ID."
// @Success  200  object  orm.Thread  "Thread Data"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
// @Resource threads
// @Route /api/v1/threads/{id} [get]
func GetOne(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get record
		var thread orm.Thread
		if err := a.DB.Where("id = ?", id).First(&thread).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusNotFound, "Thread not found")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Succeed
		response.JSON(w, http.StatusOK, thread)
	})
}
