package schedule

import "time"

type Entity struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ID        string    `db:"id"`
	StoreID   string    `db:"store_id"`
	Periods   []byte    `db:"periods" `
	IsActive  bool      `db:"is_active" `
}
