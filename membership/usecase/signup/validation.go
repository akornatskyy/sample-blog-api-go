package signup

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	e := &errorstate.ErrorState{
		Domain: "signup",
	}

	rule.Email.Validate(e, req.Email)
	rule.Username.Validate(e, req.Username)

	if rule.Password.Validate(e, req.Password) &&
		req.Password != req.ConfirmPassword {
		_ = e.Add(&errorstate.Detail{
			Type:     "field",
			Location: "password",
			Reason:   "no match",
			Message:  "Passwords do not match.",
		})
	}

	return e.OrNil()
}
