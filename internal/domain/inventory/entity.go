package inventory

import "time"

type Entity struct {
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	ID            string    `db:"id"`
	StoreID       string    `db:"store_id"`
	CatalogID     string    `db:"catalog_id"`
	ProductID     string    `db:"product_id"`
	Quantity      *string   `db:"quantity"`
	QuantityMin   *string   `db:"quantity_min"`
	QuantityMax   *string   `db:"quantity_max"`
	Price         *string   `db:"price"`
	PriceSpecial  *string   `db:"price_special"`
	PricePrevious *string   `db:"price_previous"`
	IsAvailable   *bool     `db:"is_available"`
}
