package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"warehouse-service/internal/domain/schedule"
	"warehouse-service/pkg/storage"
)

type ScheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

func (s *ScheduleRepository) Select(ctx context.Context) (dest []schedule.Entity, err error) {
	query := `
        SELECT id, store_id, periods, is_active
        FROM schedules`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *ScheduleRepository) Create(ctx context.Context, data schedule.Entity) (id string, err error) {
	query := `
        INSERT INTO schedules(store_id, periods, is_active)
        VALUES ($1, $2, $3)
        RETURNING id`

	args := []interface{}{data.StoreID, data.Periods, data.IsActive}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *ScheduleRepository) Get(ctx context.Context, storeID string) (dest *schedule.Entity, err error) {
	query := `
        SELECT id, store_id, periods, is_active
        FROM schedules
        WHERE store_id=$1`

	args := []interface{}{storeID}

	dest = new(schedule.Entity)
	if err = s.db.GetContext(ctx, dest, query, args...); err != nil {
		if err == sql.ErrNoRows {
			dest, err = nil, nil
		}
		return
	}

	return
}

func (s *ScheduleRepository) Update(ctx context.Context, storeID string, data *schedule.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, storeID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE schedules SET %s WHERE store_id=$%d", strings.Join(sets, ", "), len(args))
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

func (s *ScheduleRepository) prepareArgs(data *schedule.Entity) (sets []string, args []any) {

	if data.Periods != nil {
		args = append(args, data.Periods)
		sets = append(sets, fmt.Sprintf("periods=$%d", len(args)))
	}

	return
}

func (s *ScheduleRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM schedules
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
