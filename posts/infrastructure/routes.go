package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Routes(c *config.Config) {
	c.Router.HandlerFunc("GET", "/posts", ListPostsHandler())
	c.Router.Handle("GET", "/posts/:slug", GetPostHandler(c.Token))
	c.Router.Handle("POST", "/posts/:slug/comments", AddPostCommentHandler(c.Token))
}
