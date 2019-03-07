package config

import (
	"crypto/rand"
	"crypto/sha1"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Token  httptoken.Token
	Router *httprouter.Router
}

var (
	errNotFound = errorstate.Single(&errorstate.Detail{
		Domain:   "HTTP",
		Type:     "router",
		Location: "path",
		Reason:   "resource not found",
		Message:  "Oops! Code 404. Sorry, we can't find that resource.",
	})

	errMethodNotAllowed = errorstate.Single(&errorstate.Detail{
		Domain:   "HTTP",
		Type:     "router",
		Location: "HTTP header",
		Reason:   "method not allowed",
		Message:  "Oops! Code 405. Sorry, the HTTP method is not allowed.",
	})
)

func New() *Config {
	key := []byte(os.Getenv("KEY"))
	if len(key) == 0 {
		log.Println("WARN: using random key")
		key = make([]byte, 16)
		if _, err := rand.Read(key); err != nil {
			panic(err)
		}
	}
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpjson.Encode(w, errNotFound, http.StatusNotFound)
	})
	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpjson.Encode(w, errMethodNotAllowed, http.StatusMethodNotAllowed)
	})
	return &Config{
		Token: &httptoken.CookieToken{
			Name: "_a",
			Ticket: &ticket.Ticket{
				MaxAge: time.Duration(30) * time.Minute,
				Signer: ticket.NewSigner(sha1.New, key),
			},
		},
		Router: r,
	}
}
