package config

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akornatskyy/goext/errorstate"
	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/goext/httptoken"
	"github.com/akornatskyy/goext/security/ticket"
	"github.com/akornatskyy/sample-blog-api-go/shared/mock"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

const (
	StrategyMock = "mock"
	StrategySQL  = "sql"
)

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

type Config struct {
	DB       *sql.DB
	Router   *httprouter.Router
	Strategy string
	Token    httptoken.Token
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

	r := httprouter.New()
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpjson.Encode(w, errNotFound, http.StatusNotFound)
	})
	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpjson.Encode(w, errMethodNotAllowed, http.StatusMethodNotAllowed)
	})

	strategy := flag.String("strategy", StrategyMock, "a repository strategy")
	flag.Parse()
	log.Printf("using %s repository strategy", *strategy)

	var db *sql.DB
	var err error
	switch *strategy {
	case StrategySQL:
		s := os.Getenv("DB")
		if len(s) == 0 {
			s = "root@/sample_blog"
			log.Printf("WARN: using default DB: %s", s)
		}
		db, err = sql.Open("mysql", s)
		if err != nil {
			log.Fatalf("ERR: %s", err)
		}
	case StrategyMock:
		mock.Load("samples.json")
	}

	return &Config{
		DB:       db,
		Router:   r,
		Strategy: *strategy,
		Token: &httptoken.CookieToken{
			Name: "_a",
			Ticket: &ticket.Ticket{
				MaxAge: time.Duration(30) * time.Minute,
				Signer: ticket.NewSigner(sha1.New, key),
			},
		},
	}
}
