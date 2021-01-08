package data

type ErrorResponse struct {
	Error string `json:"error"`
}

func CustomError(err string) ErrorResponse {
	return ErrorResponse{err}
}

func ResourceNotFound() ErrorResponse {
	return ErrorResponse{"Resource not found"}
}
