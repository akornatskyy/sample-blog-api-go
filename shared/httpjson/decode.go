package httpjson

import (
	"encoding/json"
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
)

var (
	invalidContentType = errorstate.New("HTTP").Add(&errorstate.Detail{
		Type:     "header",
		Location: "Content-Type",
		Reason:   "unexpected content type",
		Message:  "Expecting 'application/json' content type.",
	})
)

// Decode reads bytes from request body as JSON-encoded
// value and stores it in the value pointed to by v.
func Decode(r *http.Request, v interface{}) error {
	if r.Header.Get("content-type") != "application/json" {
		return invalidContentType
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errorstate.New("JSON").Add(&errorstate.Detail{
			Type:     "decode",
			Location: "HTTP request body",
			Reason:   err.Error(),
			Message:  "Unable to parse JSON.",
		})
	}
	return nil
}
