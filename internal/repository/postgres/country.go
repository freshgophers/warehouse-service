package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"warehouse-service/internal/domain/country"
	"warehouse-service/pkg/storage"
)

type CountryRepository struct {
	db *sqlx.DB
}

func NewCountryRepository(db *sqlx.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (s *CountryRepository) Select(ctx context.Context) (dest []country.Entity, err error) {
	query := `
        SELECT id,name
        FROM countries`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *CountryRepository) Create(ctx context.Context, data country.Entity) (id string, err error) {
	query := `
        INSERT INTO countries (name)
        VALUES ($1)
        RETURNING id`

	args := []interface{}{data.Name}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *CountryRepository) Get(ctx context.Context, id string) (dest country.Entity, err error) {
	query := `
        SELECT id,name
        FROM countries
        WHERE id=$1`

	args := []interface{}{id}

	if err = s.db.GetContext(ctx, &dest, query, args...); err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		err = storage.ErrorNotFound
	}

	return
}

func (s *CountryRepository) Update(ctx context.Context, id string, data country.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE countries SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
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

func (s *CountryRepository) prepareArgs(data country.Entity) (sets []string, args []any) {

	if data.Name != nil {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}

	return
}

func (s *CountryRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM countries
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
