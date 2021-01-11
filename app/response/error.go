package response

import "net/http"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	})
}
