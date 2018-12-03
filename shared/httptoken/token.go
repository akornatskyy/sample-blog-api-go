package httptoken

import (
	"encoding"
	"net/http"
)

type Token interface {
	Write(w http.ResponseWriter, m encoding.BinaryMarshaler) error
	Authorize(w http.ResponseWriter, r *http.Request, p encoding.BinaryUnmarshaler) error
	Invalidate(w http.ResponseWriter)
}
