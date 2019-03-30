package search

import "github.com/akornatskyy/sample-blog-api-go/posts/domain/post"

const PageSize = 2

type (
	Request struct {
		Query string `binding:"q"`
		Page  int    `binding:"page"`
	}

	Paging struct {
		Before *int `json:"before,omitempty"`
		After  *int `json:"after,omitempty"`
	}

	Response struct {
		Paging *Paging      `json:"paging"`
		Items  []*post.Post `json:"items"`
	}
)
