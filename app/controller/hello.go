package controller

import (
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"net/http"
)

func Hello(a *core.App) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    }
}
