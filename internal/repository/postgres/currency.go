package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"

	"warehouse-service/internal/domain/currency"
	"warehouse-service/pkg/storage"
)

type CurrencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{
		db: db,
	}
}

func (s *CurrencyRepository) Select(ctx context.Context) (dest []currency.Entity, err error) {
	query := `
        SELECT id, country_id, sign, decimals,prefix
        FROM currencies`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *CurrencyRepository) Create(ctx context.Context, data currency.Entity) (id string, err error) {
	query := `
        INSERT INTO currencies (country_id, sign, decimals,prefix)
        VALUES ($1, $2, $3)
        RETURNING id`

	args := []interface{}{data.CountryID, data.Sign, data.Decimals}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *CurrencyRepository) Get(ctx context.Context, countryID string) (dest *currency.Entity, err error) {
	query := `
        SELECT id, country_id, sign, decimals,prefix
        FROM currencies
        WHERE country_id=$1`

	args := []interface{}{countryID}

	dest = new(currency.Entity)
	if err = s.db.GetContext(ctx, dest, query, args...); err != nil {
		if err == sql.ErrNoRows {
			dest, err = nil, nil
		}
		return
	}

	return
}

func (s *CurrencyRepository) Update(ctx context.Context, countryID string, data currency.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, countryID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE currencies SET %s WHERE country_id=$%d", strings.Join(sets, ", "), len(args))
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

func (s *CurrencyRepository) prepareArgs(data currency.Entity) (sets []string, args []any) {

	if data.Sign != nil {
		args = append(args, *data.Sign)
		sets = append(sets, fmt.Sprintf("sign=$%d", len(args)))
	}

	if data.Decimals != nil {
		args = append(args, *data.Decimals)
		sets = append(sets, fmt.Sprintf("decimals=$%d", len(args)))
	}

	return
}

func (s *CurrencyRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM currencies
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
