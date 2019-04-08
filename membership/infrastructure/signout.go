package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/httptoken"
)

func SignOutHandler(t httptoken.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Invalidate(w)
		w.WriteHeader(http.StatusNoContent)
	}
}
