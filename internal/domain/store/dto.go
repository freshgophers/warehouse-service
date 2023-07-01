package store

import (
	"errors"
	"github.com/shopspring/decimal"
	"net/http"
	"warehouse-service/internal/domain/city"
	"warehouse-service/internal/domain/country"
	"warehouse-service/internal/domain/currency"
	"warehouse-service/internal/domain/delivery"
	"warehouse-service/internal/domain/schedule"
)

type Request struct {
	ID         string            `json:"id"`
	MerchantID string            `json:"merchantID"`
	CityID     string            `json:"cityID"`
	Name       string            `json:"name"`
	Address    string            `json:"address"`
	Location   string            `json:"location"`
	Rating     decimal.Decimal   `json:"rating"`
	IsActive   bool              `json:"isActive"`
	City       city.Response     `json:"city"`
	Schedule   schedule.Response `json:"schedule"`
	Delivery   delivery.Response `json:"delivery"`
	Currency   currency.Response `json:"currency"`
	Area       delivery.Area     `json:"area"`
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
	ID         string             `json:"id"`
	MerchantID string             `json:"merchantID"`
	Name       string             `json:"name"`
	Address    string             `json:"address"`
	Location   string             `json:"location"`
	Rating     decimal.Decimal    `json:"rating"`
	IsActive   bool               `json:"isActive"`
	City       *city.Response     `json:"city,omitempty"`
	Country    *country.Response  `json:"country,omitempty"`
	Schedule   *schedule.Response `json:"schedule,omitempty"`
	Delivery   *delivery.Response `json:"delivery,omitempty"`
	Currency   *currency.Response `json:"currency,omitempty"`
	Area       *delivery.Area     `json:"area,omitempty"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:         data.ID,
		MerchantID: data.MerchantID,
		Name:       *data.Name,
		Address:    *data.Address,
		Location:   *data.Location,
		Rating:     *data.Rating,
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
