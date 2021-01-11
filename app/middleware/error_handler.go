package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/response"
	"net/http"
	"os"
	"path/filepath"
)

var logError *logrus.Logger

func init() {
	logError = logrus.New()
	logError.ReportCaller = true
	file, err := os.OpenFile(filepath.Join("app", "log", "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logError.Out = file
	} else {
		logError.Info("Failed to log to file, using default stderr")
	}
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logError.WithFields(logrus.Fields{
					"error": fmt.Sprintf("%v", err),
				}).Error()
				response.Error(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
