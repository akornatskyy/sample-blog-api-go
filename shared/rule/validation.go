package rule

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/validator"
)

var (
	Username = validator.String("username").Required().Min(2).Build()
	Password = validator.String("password").Required().Min(8).Build()
)
