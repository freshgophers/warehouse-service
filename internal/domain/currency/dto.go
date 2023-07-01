package currency

import (
	"errors"
	"net/http"
)

type Request struct {
	ID       string `json:"id"`
	Sign     string `json:"sing"`
	Decimals string `json:"decimals"`
	Prefix   bool   `json:"prefix"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.ID == "" {
		return errors.New("ID: cannot be blank")
	}

	if s.Sign == "" {
		return errors.New("sign: cannot be blank")
	}

	if s.Decimals == "" {
		return errors.New("decimals: cannot be blank")
	}
	return nil
}

type Response struct {
	ID       string `json:"id"`
	Sign     string `json:"sing"`
	Decimals string `json:"decimals"`
	Prefix   bool   `json:"prefix"`
}

func ParseFromEntity(data *Entity) (res *Response) {
	if data == nil {
		return
	}

	res = &Response{
		ID:       data.ID,
		Sign:     *data.Sign,
		Decimals: *data.Decimals,
		Prefix:   data.Prefix,
	}

	return
}
