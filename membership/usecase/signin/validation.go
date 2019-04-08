package signin

import (
	"github.com/akornatskyy/goext/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	e := &errorstate.ErrorState{
		Domain: "signin",
	}

	rule.Username.Validate(e, req.Username)
	rule.Password.Validate(e, req.Password)

	return e.OrNil()
}
