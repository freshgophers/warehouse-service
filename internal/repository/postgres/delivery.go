package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"warehouse-service/internal/domain/delivery"
	"warehouse-service/pkg/storage"
)

type DeliveryRepository struct {
	db *sqlx.DB
}

func NewDeliveryRepository(db *sqlx.DB) *DeliveryRepository {
	return &DeliveryRepository{
		db: db,
	}
}

func (s *DeliveryRepository) Select(ctx context.Context) (dest []delivery.Entity, err error) {
	query := `
        SELECT id, store_id, periods,areas, is_active
        FROM deliveries`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *DeliveryRepository) Create(ctx context.Context, data delivery.Entity) (id string, err error) {
	query := `
        INSERT INTO deliveries(store_id, periods,areas, is_active)
        VALUES ($1, $2, $3,$4)
        RETURNING id`

	args := []interface{}{data.StoreID, data.Periods, data.Areas, data.IsActive}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *DeliveryRepository) Get(ctx context.Context, storeID string) (dest *delivery.Entity, err error) {
	query := `
        SELECT id, store_id, periods,areas, is_active
        FROM deliveries
        WHERE store_id=$1`

	args := []interface{}{storeID}

	dest = new(delivery.Entity)
	if err = s.db.GetContext(ctx, dest, query, args...); err != nil {
		if err == sql.ErrNoRows {
			dest, err = nil, nil
		}
		return
	}

	return
}

func (s *DeliveryRepository) Update(ctx context.Context, storeID string, data *delivery.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, storeID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE deliveries SET %s WHERE store_id=$%d", strings.Join(sets, ", "), len(args))
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

func (s *DeliveryRepository) prepareArgs(data *delivery.Entity) (sets []string, args []any) {

	if data.Periods != nil {
		args = append(args, data.Periods)
		sets = append(sets, fmt.Sprintf("periods=$%d", len(args)))
	}

	if data.Areas != nil {
		args = append(args, data.Areas)
		sets = append(sets, fmt.Sprintf("areas=$%d", len(args)))
	}

	return
}

func (s *DeliveryRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM deliveries
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
