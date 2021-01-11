package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var logAccess *logrus.Logger

func init() {
	logAccess = logrus.New()
	file, err := os.OpenFile(filepath.Join("app", "log", "access.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logAccess.Out = file
	} else {
		logAccess.Info("Failed to log to file, using default stderr")
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (ww *responseWriterWrapper) WriteHeader(statusCode int) {
	ww.ResponseWriter.WriteHeader(statusCode)
	ww.statusCode = statusCode
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get current time
		start := time.Now()
		// Execute next handler
		ww := wrapResponseWriter(w)
		next.ServeHTTP(ww, r)
		// Create log
		logAccess.WithFields(logrus.Fields{
			"method":   r.Method,
			"uri":      r.RequestURI,
			"ip":       r.RemoteAddr,
			"duration": time.Since(start),
			"code":     ww.statusCode,
		}).Info()
	})
}
