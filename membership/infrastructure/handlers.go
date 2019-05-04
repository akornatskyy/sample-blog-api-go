package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/errorstate"
	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/goext/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/usecase/signin"
	"github.com/akornatskyy/sample-blog-api-go/membership/usecase/signup"
	"github.com/akornatskyy/sample-blog-api-go/shared/security"
)

func SignInHandler(t httptoken.Token) http.HandlerFunc {
	type response struct {
		Username string `json:"username"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
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
		p := security.Principal{ID: resp.UserID}
		if t.Write(w, &p) != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		httpjson.Encode(w, &response{Username: req.Username}, http.StatusOK)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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

func SignOutHandler(t httptoken.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Invalidate(w)
		w.WriteHeader(http.StatusNoContent)
	}
}

var (
	errForbidden = errorstate.Single(&errorstate.Detail{
		Domain:   "HTTP",
		Type:     "authorization",
		Location: "token",
		Reason:   "forbidden",
		Message: "You do not have permission to access this resource using " +
			"the credentials that you supplied.",
	})
)

func UserHandler(t httptoken.Token) http.HandlerFunc {
	type response struct {
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var p security.Principal
		if err := t.Authorize(w, r, &p); err != nil {
			httpjson.Encode(w, errForbidden, http.StatusForbidden)
			return
		}
		u, err := domain.UserRepository().FindUserByID(p.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		resp := response{
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
		httpjson.Encode(w, &resp, http.StatusOK)
	}
}
