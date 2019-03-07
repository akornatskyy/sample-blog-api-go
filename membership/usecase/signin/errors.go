package signin

import "github.com/akornatskyy/sample-blog-api-go/shared/errorstate"

var (
	errInvalidCredentials = errorstate.Single(&errorstate.Detail{
		Domain:   "signin",
		Type:     "summary",
		Location: "user",
		Reason:   "invalid credentials",
		Message:  "The username or password provided is incorrect.",
	})

	errUserLocked = errorstate.Single(&errorstate.Detail{
		Domain:   "signin",
		Type:     "summary",
		Location: "user",
		Reason:   "account locked",
		Message:  "The account is locked. Contact system administrator, please.",
	})
)
