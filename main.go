package main

import (
	"log"
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

const (
	addr = ":8080"
)

func main() {
	c := config.New()

	membership.Setup(c)

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, c.Router))
}
