package ticket

import (
	"crypto/hmac"
	"encoding/base64"
	"errors"
	"hash"
)

var (
	errInvalidLength     = errors.New("invalid length")
	errSignatureMismatch = errors.New("signature mismatch")
)

type Signer struct {
	h   func() hash.Hash
	key []byte
}

func NewSigner(h func() hash.Hash, key []byte) *Signer {
	return &Signer{
		h:   h,
		key: key,
	}
}

func (s *Signer) EncodeToString(src []byte) (string, error) {
	h := hmac.New(s.h, s.key)
	if _, err := h.Write(src); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(h.Sum(src)), nil
}

func (t *Signer) DecodeString(s string) ([]byte, error) {
	b, _ := base64.RawURLEncoding.DecodeString(s)
	h := hmac.New(t.h, t.key)
	n := len(b) - h.Size()
	if n < 0 {
		return nil, errInvalidLength
	}
	value := b[:n]
	if _, err := h.Write(value); err != nil {
		return nil, err
	}
	if !hmac.Equal(h.Sum(nil), b[n:]) {
		return nil, errSignatureMismatch
	}
	return value, nil
}
