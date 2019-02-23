package config

import (
	"crypto/rand"
	"crypto/sha1"
	"log"
	"os"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
)

type Config struct {
	Token httptoken.Token
}

func New() *Config {
	key := []byte(os.Getenv("KEY"))
	if len(key) == 0 {
		log.Println("WARN: using random key")
		key = make([]byte, 16)
		if _, err := rand.Read(key); err != nil {
			panic(err)
		}
	}
	return &Config{
		Token: &httptoken.CookieToken{
			Name: "_a",
			Ticket: &ticket.Ticket{
				MaxAge: time.Duration(30) * time.Minute,
				Signer: ticket.NewSigner(sha1.New, key),
			},
		},
	}
}
