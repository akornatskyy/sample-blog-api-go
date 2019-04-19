package addcomment

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/security"
)

type (
	Request struct {
		Slug      string
		Message   string
		Principal security.Principal
	}

	Response struct {
	}
)
