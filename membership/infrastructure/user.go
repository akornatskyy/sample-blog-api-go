package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security"
)

var (
	ErrForbidden = errorstate.New("HTTP").Add(&errorstate.Detail{
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
		p := security.Principal{}
		if err := t.Authorize(w, r, &p); err != nil {
			httpjson.Encode(w, ErrForbidden, http.StatusForbidden)
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
