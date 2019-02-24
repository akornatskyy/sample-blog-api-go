package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Routes(c *config.Config) {
	c.Router.HandlerFunc("POST", "/signin", SignInHandler(c.Token))
	c.Router.HandlerFunc("POST", "/signup", SignUpHandler)
	c.Router.HandlerFunc("GET", "/signout", SignOutHandler(c.Token))
	c.Router.HandlerFunc("GET", "/user", UserHandler(c.Token))
}
