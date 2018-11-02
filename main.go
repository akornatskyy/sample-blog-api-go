package main

import (
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

func main() {
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
