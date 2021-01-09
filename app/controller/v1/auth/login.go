package auth

import (
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"net/http"
)

func Login(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        response.JSON(w, http.StatusNotImplemented, "Not Implemented")
    })
}
