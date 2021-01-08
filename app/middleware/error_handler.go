package middleware

import (
	"github.com/naufal-dean/onboarding-dean-local/app/model/data"
	"github.com/naufal-dean/onboarding-dean-local/app/response"
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				response.JSON(w, http.StatusInternalServerError, data.InternalServerError())
			}
		}()

		next.ServeHTTP(w, r)
	})
}
