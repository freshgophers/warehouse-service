package store

import (
	"errors"
	"github.com/shopspring/decimal"
	"net/http"
)

type Request struct {
	MerchantID string          `json:"merchantID"`
	CityID     string          `json:"cityID"`
	Name       string          `json:"name"`
	Address    string          `json:"address"`
	Location   string          `json:"location"`
	Rating     decimal.Decimal `json:"rating"`
	IsActive   bool            `json:"isActive"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.MerchantID == "" {
		return errors.New("merchantID: cannot be blank")
	}

	if s.CityID == "" {
		return errors.New("cityID: cannot be blank")
	}

	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}

	if s.Address == "" {
		return errors.New("address: cannot be blank")
	}

	if s.Location == "" {
		return errors.New("location: cannot be blank")
	}

	return nil
}

type Response struct {
	ID         string          `json:"id"`
	MerchantID string          `json:"merchantID"`
	Name       string          `json:"name"`
	Location   string          `json:"location"`
	Rating     decimal.Decimal `json:"rating"`
	City       City            `json:"city"`
	Schedule   Schedule        `json:"schedule"`
	Delivery   Delivery        `json:"delivery"`
	Areas      Areas           `json:"areas"`
	Address    string          `json:"address"`
	IsActive   bool            `json:"isActive"`
}

type City struct {
	ID        string `json:"ID" db:"id"`
	CountryID string `json:"countryID" db:"country_id"`
	Name      string `json:"name" db:"name"`
	Geocenter string `json:"geocenter" db:"geocenter"`
}

type Schedule struct {
	IsActive bool    `json:"isActive" `
	Periods  Periods `json:"periods" `
}

type Periods struct {
	Day  string `json:"day" `
	From string `json:"from" `
	To   string `json:"to" `
}

type Delivery struct {
	IsActive bool    `json:"isActive" `
	Periods  Periods `json:"periods" `
}

type Areas struct {
	X string `json:"x" `
	Y string `json:"y" `
}

type Currency struct {
	ID      string `json:"id"`
	Sign    string `json:"sign" `
	Decimal string `json:"decimal"`
	Prefix  bool   `json:"prefix" `
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:         data.ID,
		MerchantID: data.MerchantID,
		Name:       *data.Name,
		Address:    *data.Address,
		Location:   *data.Location,
		Rating:     *data.Rating,
		IsActive:   *data.IsActive,
	}
	return
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
