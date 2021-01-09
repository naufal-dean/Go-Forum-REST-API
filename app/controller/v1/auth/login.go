package auth

import (
	"encoding/json"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/lib/auth"
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(a *core.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get input
		var input LoginInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, data.CustomError("Invalid input"))
			return
		}
		defer r.Body.Close()

		// Check username and password
		// TODO: use bcrypt
		var user orm.User
		if err := a.DB.Where("email = ?", input.Email).Where("password = ?", input.Password).First(&user).Error;
			err != nil {
			// TODO: create response object
			response.JSON(w, http.StatusNotFound, "Invalid email or password")
			return
		}

		// Create jwt token
		token, claims := auth.NewToken(user.ID)
		if err := a.DB.Create(&orm.Token{UserID: claims.UserID, TokenUUID: claims.TokenUUID}).Error;
			err != nil {
			response.JSON(w, http.StatusInternalServerError, data.CustomError("Login failed"))
			return
		}

		// Succeed
		response.JSON(w, http.StatusOK, token)
	})
}
