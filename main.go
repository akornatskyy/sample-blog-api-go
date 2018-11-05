package main

import (
	"log"
	"net/http"

	m "github.com/akornatskyy/sample-blog-api-go/membership/infrastructure"
)

const (
	addr = ":8080"
)

func main() {
	http.HandleFunc("/signin", m.SignInHandler)

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
