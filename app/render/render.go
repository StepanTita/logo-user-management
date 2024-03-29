package render

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

// Message, function to render json response message
func Message(message interface{}) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

// Valid json respond rendering
func Respond(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	if status == http.StatusNoContent {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}
