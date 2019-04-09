package getpost

import (
	"github.com/akornatskyy/goext/errorstate"
)

var (
	ErrNotFound = errorstate.Single(&errorstate.Detail{
		Domain:   "posts",
		Type:     "summary",
		Location: "post",
		Reason:   "not found",
		Message:  "The post cannot be found.",
	})
)
