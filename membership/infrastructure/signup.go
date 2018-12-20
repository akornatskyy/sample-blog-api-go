package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership/usecase/signup"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req signup.Request
	if err := httpjson.Decode(r, &req, 140); err != nil {
		httpjson.Encode(w, err, http.StatusUnprocessableEntity)
		return
	}
	if _, err := signup.Process(&req); err != nil {
		httpjson.Encode(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
