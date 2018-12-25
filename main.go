package main

import (
	"crypto/rand"
	"crypto/sha1"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/membership"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
)

const (
	addr = ":8080"
)

func main() {
	key := []byte(os.Getenv("KEY"))
	if len(key) == 0 {
		log.Println("WARN: using random key")
		key = make([]byte, 16)
		if _, err := rand.Read(key); err != nil {
			panic(err)
		}
	}
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
