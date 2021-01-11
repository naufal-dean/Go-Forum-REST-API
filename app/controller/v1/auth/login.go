package auth

import (
	"encoding/json"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

// @Title Login.
// @Description Login with email and password.
// @Param  user  body  LoginInput  true  "New user data."
// @Success  200  array  LoginSuccessResponse  "Access Token JSON"
// @Failure  401  object  response.ErrorResponse  "Invalid Credentials Error JSON"
// @Failure  422  object  response.ErrorResponse  "Invalid Input Error JSON"
// @Resource auth
// @Route /api/v1/login [post]
func Login(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input LoginInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Invalid input")
			return
		}
		defer r.Body.Close()

		// Check username and password
		// TODO: use bcrypt
		var user orm.User
		if err := a.DB.Where("email = ?", input.Email).Where("password = ?", input.Password).First(&user).Error;
			err != nil {
			// TODO: create response object
			response.Error(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		// Create jwt token
		token, claims := auth.NewToken(user.ID)
		if err := a.DB.Create(&orm.Token{UserID: claims.UserID, TokenUUID: claims.TokenUUID}).Error;
			err != nil {
			response.Error(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Succeed
		response.JSON(w, http.StatusOK, LoginSuccessResponse{
			TokenType: "bearer",
			Token:     token,
		})
	})
}
