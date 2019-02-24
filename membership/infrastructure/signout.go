package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
)

func SignOutHandler(t httptoken.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Invalidate(w)
	}
}
