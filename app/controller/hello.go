package controller

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"net/http"
)

func Hello(a *core.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
    })
}
