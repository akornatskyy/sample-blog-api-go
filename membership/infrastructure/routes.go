package infrastructure

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/signin", SignInHandler)
}
