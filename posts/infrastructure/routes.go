package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Routes(c *config.Config) {
	c.Router.HandlerFunc("GET", "/posts", ListPostsHandler())
	c.Router.Handle("GET", "/posts/:id", GetPostHandler())
	c.Router.Handle("GET", "/posts/:id/comments", ListPostCommentsHandler())
}
