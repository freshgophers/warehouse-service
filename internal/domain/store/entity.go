package store

import (
	"github.com/shopspring/decimal"
	"time"
)

type Entity struct {
	CreatedAt  time.Time        `db:"created_at"`
	UpdatedAt  time.Time        `db:"updated_at"`
	ID         string           `db:"id"`
	MerchantID string           `db:"merchant_id"`
	CityID     string           `db:"city_id"`
	Name       *string          `db:"name"`
	Address    *string          `db:"address"`
	Location   *string          `db:"location"`
	Rating     *decimal.Decimal `db:"rating"`
	IsActive   *bool            `db:"is_active"`
}
