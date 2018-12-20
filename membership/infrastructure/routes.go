package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
)

func Routes(t httptoken.Token) {
	http.HandleFunc("/signin", SignInHandler(t))
	http.HandleFunc("/signup", SignUpHandler)
}
