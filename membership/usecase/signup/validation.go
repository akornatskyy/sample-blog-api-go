package signup

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	err := errorstate.New("signup")

	rule.Email.Validate(err, req.Email)
	rule.Username.Validate(err, req.Username)

	if rule.Password.Validate(err, req.Password) &&
		req.Password != req.ConfirmPassword {
		_ = err.Add(&errorstate.Detail{
			Type:     "field",
			Location: "password",
			Reason:   "no match",
			Message:  "Passwords do not match.",
		})
	}

	return err.OrNil()
}
