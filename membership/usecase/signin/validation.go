package signin

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	err := errorstate.New("signin")

	rule.Username.Validate(err, req.Username)
	rule.Password.Validate(err, req.Password)

	return err.OrNil()
}
