package ticket

import (
	"encoding/binary"
	"errors"
	"time"
)

type Ticket struct {
	MaxAge time.Duration
	Signer *Signer
}

var (
	ErrExpired = errors.New("expired")
)

func (t *Ticket) EncodeToString(b []byte) (string, error) {
	buf := make([]byte, binary.MaxVarintLen64)
	expires := time.Now().Add(t.MaxAge).Unix()
	n := binary.PutVarint(buf, expires)
	return t.Signer.EncodeToString(append(buf[:n], b...))
}

func (t *Ticket) DecodeString(s string) ([]byte, int, error) {
	b, err := t.Signer.DecodeString(s)
	if err != nil {
		return nil, 0, err
	}
	expires, n := binary.Varint(b)
	remaining := int(expires - time.Now().Unix())
	if remaining <= 0 {
		return nil, remaining, ErrExpired
	}
	return b[n:], remaining, nil
}
