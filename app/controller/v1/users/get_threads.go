package users

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

// @Title Get user's thread.
// @Description Get all threads owned by user.
// @Param  id  path  int  true  "User ID."
// @Success  200  array  []orm.Thread  "Array of Thread's Thread"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
// @Resource users/threads
// @Route /api/v1/users/{id}/threads [get]
func GetThreads(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get records
		var user orm.User
		if err := a.DB.Preload("Threads").Where("id = ?", id).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(w, http.StatusNotFound, "User not found")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Succeed
		response.JSON(w, http.StatusOK, user.Threads)
	})
}
