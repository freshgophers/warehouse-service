package city

import "time"

type Entity struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ID        string    `db:"id"`
	CountryID string    `db:"country_id"`
	Name      *string   `db:"name"`
	GeoCenter *string   `db:"geocenter"`
}
