package httpjson

import (
	"encoding/json"
	"net/http"
)

// Encode writes the JSON encoding of v to the response writer.
func Encode(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
