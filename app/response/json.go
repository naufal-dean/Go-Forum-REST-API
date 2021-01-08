package response

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(errors.Wrap(err, "JSON response error"))
		}
	}
}
