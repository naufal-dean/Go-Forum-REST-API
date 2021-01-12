package auth

import (
	"encoding/json"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/cerror"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response/data"
	"net/http"
)

type RegisterInput struct {
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Name                 string `json:"name" validate:"required"`
}

// @Title Register.
// @Description Register a new user account.
// @Param  user  body  RegisterInput  true  "New user data."
// @Success  201  array  response.SuccessResponse  "Register Succeed JSON"
// @Failure  422  object  response.ErrorResponse  "Invalid Input Error JSON"
// @Resource auth
// @Route /api/v1/register [post]
func Register(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input RegisterInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

		// Validate input
		err := a.Validate.Struct(input)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, data.NewValidationErrorResponse(err))
			return
		}

		// Create record
		user := orm.User{Email: input.Email, Password: input.Password, Name: input.Name}
		if err := a.DB.Create(&user).Error; err != nil {
			if a.DB.Where("email = ?", input.Email).First(&orm.User{}).Error == nil {
				response.Error(w, http.StatusConflict, "Email is already in use")
				return
			} else {
				panic(&cerror.DatabaseError{DBErr: err})
			}
		}

		// Succeed
		response.Success(w, http.StatusCreated, "Register succeed")
	})
}
