package main

import (
	"log"
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership"
)

const (
	addr = ":8080"
)

func main() {
	membership.Setup()

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
