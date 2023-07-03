package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"

	"warehouse-service/internal/domain/city"
	"warehouse-service/pkg/storage"
)

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository(db *sqlx.DB) *CityRepository {
	return &CityRepository{
		db: db,
	}
}

func (s *CityRepository) Select(ctx context.Context) (dest []city.Entity, err error) {
	query := `
        SELECT id, country_id, name, geocenter
        FROM cities`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *CityRepository) Create(ctx context.Context, data city.Entity) (id string, err error) {
	query := `
        INSERT INTO cities (country_id, name, geocenter)
        VALUES ($1, $2, $3)
        RETURNING id`

	args := []interface{}{data.CountryID, data.Name, data.GeoCenter}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *CityRepository) Get(ctx context.Context, id string) (dest *city.Entity, err error) {
	query := `
        SELECT id, country_id, name, geocenter
        FROM cities
        WHERE id=$1`

	args := []interface{}{id}

	dest = new(city.Entity)
	if err = s.db.GetContext(ctx, dest, query, args...); err != nil {
		if err == sql.ErrNoRows {
			dest, err = nil, nil
		}
		return
	}

	return
}

func (s *CityRepository) Update(ctx context.Context, id string, data city.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE cities SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return
		}

		if err == sql.ErrNoRows {
			err = storage.ErrorNotFound
		}
	}

	return
}

func (s *CityRepository) prepareArgs(data city.Entity) (sets []string, args []interface{}) {
	if data.Name != nil {
		args = append(args, *data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}

	if data.GeoCenter != nil {
		args = append(args, *data.GeoCenter)
		sets = append(sets, fmt.Sprintf("geocenter=$%d", len(args)))
	}

	return
}

func (s *CityRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM cities
        WHERE id=$1`

	args := []interface{}{id}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		err = storage.ErrorNotFound
	}

	return
}
