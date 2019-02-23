package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Routes(c *config.Config) {
	http.HandleFunc("/signin", SignInHandler(c.Token))
	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/signout", SignOutHandler(c.Token))
	http.HandleFunc("/user", UserHandler(c.Token))
}
