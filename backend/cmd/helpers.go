package cmd

import (
	"encoding/json"
	"net/http"
)

// setHeaders sets the required headers for response.
func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

// errorJSON returns an api message with the error as json.
func errorJSON(w http.ResponseWriter, error string, code int) {
	setHeaders(w)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiMessage{Success: false, Message: error})
}
