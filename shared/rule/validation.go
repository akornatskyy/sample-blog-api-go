package rule

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/validator"
)

var (
	Username = validator.String("username").Required().Min(2).Max(20).Build()
	Password = validator.String("password").Required().Min(8).Max(12).Build()
)
