package currency

import "time"

type Entity struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ID        string    `db:"id"`
	CountryID string    `db:"country_id"`
	Sign      *string   `db:"sign"`
	Decimals  *string   `db:"decimals"`
	Prefix    bool      `db:"prefix"`
}
