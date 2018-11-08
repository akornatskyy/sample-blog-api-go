package httpjson

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Decode reads bytes from request body as JSON-encoded
// value and stores it in the value pointed to by v.
func Decode(r *http.Request, v interface{}) error {
	if r.Header.Get("content-type") != "application/json" {
		return errors.New("Expecting 'application/json' content type.")
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.New("Unable to parse JSON.")
	}
	return nil
}
