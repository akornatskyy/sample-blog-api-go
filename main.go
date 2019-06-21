// A simple blog API.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/akornatskyy/sample-blog-api-go/membership"
	"github.com/akornatskyy/sample-blog-api-go/posts"
	"github.com/akornatskyy/sample-blog-api-go/public"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

const (
	addr = ":8080"
)

func main() {
	log.Printf("starting...")
	c := config.New()

	membership.Setup(c)
	posts.Setup(c)
	public.Setup(c)

	log.Printf("http listening on %s", addr)
	server := &http.Server{Addr: addr, Handler: c.Router}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Println("http stopped")
			default:
				log.Printf("http serve: %s", err)
			}
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown: %s", err)
	}
	if err := server.Close(); err != nil {
		log.Printf("server close: %s", err)
	}
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			log.Printf("db close: %s", err)
		}
	}

	log.Println("done")
}
