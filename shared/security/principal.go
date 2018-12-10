package security

import (
	"strconv"
)

type Principal struct {
	ID int
}

func (p *Principal) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(p.ID)), nil
}

func (p *Principal) UnmarshalBinary(data []byte) error {
	id, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}
