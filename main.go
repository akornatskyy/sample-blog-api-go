package main

import (
	"log"
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/membership"
	m "github.com/akornatskyy/sample-blog-api-go/membership/infrastructure"
)

const (
	addr = ":8080"
)

func main() {
	membership.Setup()
	http.HandleFunc("/signin", m.SignInHandler)

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
