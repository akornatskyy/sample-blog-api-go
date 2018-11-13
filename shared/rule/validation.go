package rule

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/validator"
)

var (
	Username = validator.String("username").Required().Build()
	Password = validator.String("password").Required().Build()
)
