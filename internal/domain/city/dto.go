package city

import (
	"errors"
	"net/http"
)

type Request struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GeoCenter string `json:"geocenter"`
	CountryID string `json:"country_id"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.ID == "" {
		return errors.New("ID: cannot be blank")
	}

	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}

	if s.GeoCenter == "" {
		return errors.New("geocenter: cannot be blank")
	}

	return nil
}

type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GeoCenter string `json:"geocenter"`
	CountryID string `json:"country_id"`
}

func ParseFromEntity(data *Entity) (res *Response) {
	if data == nil {
		return
	}

	res = &Response{
		ID:        data.ID,
		Name:      *data.Name,
		GeoCenter: *data.GeoCenter,
		CountryID: data.CountryID,
	}

	return
}

func ParseFromEntities(data []*Entity) (res []*Response) {
	res = make([]*Response, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
