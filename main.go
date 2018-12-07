package main

import (
	"crypto/sha1"
	"log"
	"net/http"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/membership"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
)

const (
	addr = ":8080"
)

func main() {
	key := []byte("secret")
	t := &httptoken.CookieToken{
		Name: "_a",
		Ticket: &ticket.Ticket{
			MaxAge: time.Duration(30) * time.Minute,
			Signer: ticket.NewSigner(sha1.New, key),
		},
	}

	membership.Setup(t)

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
