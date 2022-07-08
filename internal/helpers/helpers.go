package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/yalagtyarzh/rest/internal/models"
)

// responseError provides error response struct
type responseError struct {
	Error string `json:"error"`
}

// ThrowError throws a responseError with the passed response status and error
func ThrowError(w http.ResponseWriter, status int, err error) {
	resp := responseError{
		Error: err.Error(),
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}

// responseOK is a valid json response struct
type responseOK struct {
	Message string        `json:"message"`
	Events  []models.User `json:"users"`
}

// WriteResponse throws a responseOK with the passed status, message, and users' slice
func WriteResponse(w http.ResponseWriter, status int, msg string, events []models.User) {
	resp := responseOK{
		Message: msg,
		Events:  events,
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}
