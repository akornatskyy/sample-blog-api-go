package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	httpjson.Encode(w, &struct{}{}, http.StatusOK)
}
