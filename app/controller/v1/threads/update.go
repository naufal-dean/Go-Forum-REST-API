package threads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/auth"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/util"
	"github.com/naufal-dean/go-forum-rest-api/app/model/cerror"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
	"gorm.io/gorm"
	"net/http"
)

type UpdateInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// @Title Update a thread.
// @Description Update a thread with ID.
// @Param  id  path  int  true  "Thread ID."
// @Param  thread  body  UpdateInput  optional  "Thread new data."
// @Success  200  object  orm.Post  "Updated Thread Data"
// @Failure  401  object  response.ErrorResponse  "Unauthorized Error"
// @Failure  403  object  response.ErrorResponse  "Forbidden Error (Non Owner)"
// @Failure  404  object  response.ErrorResponse  "Resource Not Found Error"
// @Failure  422  object  response.ErrorResponse  "Unprocessable Input Error"
// @Resource threads
// @Route /api/v1/threads/{id} [put]
func Update(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params
		id, err := util.StrToUint(mux.Vars(r)["id"])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid id")
			return
		}

		// Get input
		var input UpdateInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

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

		// Check ownership
		claims, ok := r.Context().Value("claims").(*auth.TokenClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "No token claims found")
			return
		}
		if claims.UserID != thread.UserID {
			response.Error(w, http.StatusForbidden, "You are not the owner of the thread")
			return
		}

		// Update record
		a.DB.Model(&thread).Updates(orm.Thread{Name: input.Name, Description: input.Description})

		// Succeed
		response.JSON(w, http.StatusOK, thread)
	})
}
