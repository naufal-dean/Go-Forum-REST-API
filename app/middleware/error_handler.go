package middleware

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/naufal-dean/go-forum-rest-api/app/model/cerror"
	"github.com/naufal-dean/go-forum-rest-api/app/response"
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
			if obj := recover(); obj != nil {
				// Log error
				logError.WithFields(logrus.Fields{
					"error": fmt.Sprintf("%v", obj),
				}).Error()
				fmt.Printf("%+v\n", obj)
				// Create error response
				err, _ := obj.(error)
				switch errors.Cause(err).(type) {
				case *cerror.DatabaseError:
					response.Error(w, http.StatusInternalServerError, "Database Error")
				default:
					response.Error(w, http.StatusInternalServerError, "Internal Server Error")
				}
			}
		}()
		// Serve next handler
		next.ServeHTTP(w, r)
	})
}
