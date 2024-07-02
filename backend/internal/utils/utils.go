package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSON encodes the given struct as JSON and writes it to the response writer.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// WriteError writes an error message as a JSON response with the given status code.
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// ParseJSON decodes a JSON request body into the given struct.
func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
