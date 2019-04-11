package security

type Principal struct {
	ID string
}

func (p *Principal) MarshalBinary() (data []byte, err error) {
	return []byte(p.ID), nil
}

func (p *Principal) UnmarshalBinary(data []byte) error {
	p.ID = string(data)
	return nil
}

func (p *Principal) IsAuthenticated() bool {
	return p.ID != ""
}
