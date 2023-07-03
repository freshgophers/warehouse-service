package inventory

import (
	"errors"
	"net/http"
)

type Request struct {
	ID            string `json:"id"`
	StoreID       string `json:"store_id"`
	ProductID     string `json:"product_id"`
	Quantity      string `json:"quantity"`
	QuantityMin   string `json:"quantity_min"`
	QuantityMax   string `json:"quantity_max"`
	Price         string `json:"price"`
	PriceSpecial  string `json:"price_special"`
	PricePrevious string `json:"price_previous"`
	IsAvailable   string `json:"is_available"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.StoreID == "" {
		return errors.New("store_id: cannot be blank")
	}

	if s.ProductID == "" {
		return errors.New("product_id: cannot be blank")
	}

	if s.Quantity == "" {
		return errors.New("name: quantity be blank")
	}

	if s.Price == "" {
		return errors.New("address: price be blank")
	}

	return nil
}

type Response struct {
	ID            string `json:"id"`
	StoreID       string `json:"store_id"`
	ProductID     string `json:"product_id"`
	Quantity      string `json:"quantity"`
	QuantityMin   string `json:"quantity_min"`
	QuantityMax   string `json:"quantity_max"`
	Price         string `json:"price"`
	PriceSpecial  string `json:"price_special"`
	PricePrevious string `json:"price_previous"`
	IsAvailable   string `json:"is_available"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:            data.ID,
		StoreID:       data.StoreID,
		Quantity:      *data.Quantity,
		QuantityMin:   *data.QuantityMin,
		QuantityMax:   *data.QuantityMax,
		Price:         *data.Price,
		PriceSpecial:  *data.PriceSpecial,
		PricePrevious: *data.PricePrevious,
		IsAvailable:   *data.IsAvailable,
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
