package ticket_test

import (
	"crypto/sha1"
	"fmt"

	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
)

var key = []byte("secret")

func ExampleSigner_EncodeToString() {
	s := ticket.NewSigner(sha1.New, key)

	fmt.Println(s.EncodeToString([]byte("test")))
	// Output: dGVzdBqjSVhe1-y9O5xIajAGfjlcpLNW <nil>
}

func ExampleSigner_DecodeString() {
	s := ticket.NewSigner(sha1.New, key)

	text, err := s.DecodeString("dGVzdBqjSVhe1-y9O5xIajAGfjlcpLNW")
	fmt.Println(string(text), err)
	// Output: test <nil>
}
