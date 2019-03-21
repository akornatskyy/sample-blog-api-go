package search

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp := Response{}
	return &resp, nil
}
