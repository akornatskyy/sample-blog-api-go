package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership/usecase/signin"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req signin.Request
	if err := httpjson.Decode(r, &req, 128); err != nil {
		httpjson.Encode(w, err, http.StatusUnprocessableEntity)
		return
	}
	resp := &signin.Response{
		Username: req.Username,
	}
	httpjson.Encode(w, resp, http.StatusOK)
}
