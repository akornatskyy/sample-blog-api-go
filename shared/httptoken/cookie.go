package httptoken

import (
	"encoding"
	"net/http"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/shared/security/ticket"
)

type CookieToken struct {
	Name   string
	Ticket *ticket.Ticket
}

func (t *CookieToken) Write(w http.ResponseWriter, p encoding.BinaryMarshaler) error {
	b, err := p.MarshalBinary()
	if err != nil {
		return err
	}
	token, err := t.Ticket.EncodeToString(b)
	if err != nil {
		return err
	}
	http.SetCookie(w, t.cookie(token))
	return nil
}

func (t *CookieToken) Authorize(w http.ResponseWriter, r *http.Request, p encoding.BinaryUnmarshaler) error {
	c, err := r.Cookie(t.Name)
	if err != nil {
		return err
	}
	b, remaining, err := t.Ticket.DecodeString(c.Value)
	if err != nil {
		return err
	}
	if remaining < int(t.Ticket.MaxAge/time.Second)/2 {
		token, err := t.Ticket.EncodeToString(b)
		if err != nil {
			return err
		}
		http.SetCookie(w, t.cookie(token))
	}
	if p == nil {
		return nil
	}
	return p.UnmarshalBinary(b)
}

func (t *CookieToken) Invalidate(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    t.Name,
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
}

func (t *CookieToken) cookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:     t.Name,
		Path:     "/",
		Expires:  time.Now().Add(t.Ticket.MaxAge),
		HttpOnly: true,
		Value:    value,
	}
}
