package security

type Principal struct {
	Username string
}

func (p *Principal) MarshalBinary() (data []byte, err error) {
	return []byte(p.Username), nil
}

func (p *Principal) UnmarshalBinary(data []byte) error {
	p.Username = string(data)
	return nil
}
