package auth

import (
	"encoding/json"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/data"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

type RegisterInput struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Name                 string `json:"name"`
}

// TODO: create succeed response

// @Title Register.
// @Description Register a new user account.
// @Param  user  body  RegisterInput  true  "New user data."
// @Success  200  array  []orm.Thread  "Register Succeed JSON"
// @Failure  422  object  data.ErrorResponse  "Invalid Input Error JSON"
// @Resource auth
// @Route /api/v1/register [post]
func Register(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input RegisterInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, data.CustomError("Invalid input"))
			return
		}
		defer r.Body.Close()

		// Check if password confirmation match
		if input.Password != input.PasswordConfirmation {
			response.JSON(w, http.StatusUnprocessableEntity, data.CustomError("Password confirmation mismatch"))
			return
		}

		// Create record
		user := orm.User{Email: input.Email, Password: input.Password, Name: input.Name}
		if err := a.DB.Create(&user).Error; err != nil {
			// TODO: create not unique error message
			response.JSON(w, http.StatusInternalServerError, data.CustomError("Create user failed"))
			return
		}

		// Succeed
		// TODO: create response object
		response.JSON(w, http.StatusCreated, "Register succeed")
	})
}
