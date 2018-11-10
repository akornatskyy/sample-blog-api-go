package httpjson

import (
	"encoding/json"
	"fmt"
	"io"
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

// Decode reads n bytes from request body as JSON-encoded
// value and stores it in the value pointed to by v.
func Decode(r *http.Request, v interface{}, n int64) error {
	if r.Header.Get("content-type") != "application/json" {
		return invalidContentType
	}
	reader := &io.LimitedReader{R: r.Body, N: n}
	if err := json.NewDecoder(reader).Decode(v); err != nil {
		if reader.N == 0 {
			return errorstate.New("HTTP").Add(&errorstate.Detail{
				Type:     "reader",
				Location: "HTTP request body",
				Reason:   "request entity too large",
				Message:  fmt.Sprintf("Request body size is limited to %d bytes.", n),
			})
		}
		return errorstate.New("JSON").Add(&errorstate.Detail{
			Type:     "decode",
			Location: "HTTP request body",
			Reason:   err.Error(),
			Message:  "Unable to parse JSON.",
		})
	}
	return nil
}
