package country

import (
	"errors"
	"net/http"
)

type Request struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.ID == "" {
		return errors.New("ID: cannot be blank")
	}

	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}

	return nil
}

type Response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ParseFromEntity(data *Entity) (res *Response) {
	if data == nil {
		return
	}

	res = &Response{
		ID:   data.ID,
		Name: *data.Name,
	}

	return
}
