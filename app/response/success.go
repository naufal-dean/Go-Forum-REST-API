package response

import "net/http"

type SuccessResponse struct {
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, SuccessResponse{
		Message: message,
	})
}
