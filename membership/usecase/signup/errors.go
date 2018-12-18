package signup

import "github.com/akornatskyy/sample-blog-api-go/shared/errorstate"

var (
	errUsernameTaken = errorstate.New("signup").Add(&errorstate.Detail{
		Type:     "field",
		Location: "username",
		Reason:   "username taken",
		Message:  "The user with such username is already registered. Please try another.",
	})

	errCreateFailed = errorstate.New("signup").Add(&errorstate.Detail{
		Type:     "summary",
		Location: "user",
		Reason:   "create failed",
		Message:  "The system was unable to create an account for you. Please try again later.",
	})
)
