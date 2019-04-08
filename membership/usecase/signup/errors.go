package signup

import (
	"github.com/akornatskyy/goext/errorstate"
)

var (
	errUsernameTaken = errorstate.Single(&errorstate.Detail{
		Domain:   "signup",
		Type:     "field",
		Location: "username",
		Reason:   "username taken",
		Message:  "The user with such username is already registered. Please try another.",
	})

	errCreateFailed = errorstate.Single(&errorstate.Detail{
		Domain:   "signup",
		Type:     "summary",
		Location: "user",
		Reason:   "create failed",
		Message:  "The system was unable to create an account for you. Please try again later.",
	})
)
