package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership/usecase/signin"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security"
)

func SignInHandler(t httptoken.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req signin.Request
		if err := httpjson.Decode(r, &req, 128); err != nil {
			httpjson.Encode(w, err, http.StatusUnprocessableEntity)
			return
		}
		resp, err := signin.Process(&req)
		if err != nil {
			httpjson.Encode(w, err, http.StatusBadRequest)
			return
		}
		p := security.Principal{Username: resp.Username}
		if t.Write(w, &p) != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		httpjson.Encode(w, resp, http.StatusOK)
	}
}
