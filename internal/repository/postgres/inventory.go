package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"

	"warehouse-service/internal/domain/inventory"
	"warehouse-service/pkg/storage"
)

type InventoryRepository struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (s *InventoryRepository) Select(ctx context.Context) (dest []inventory.Entity, err error) {
	query := `
        SELECT id, store_id,product_id quantity, quantity_min,quantity_max,price,price_sepcial,price_previous,is_available
        FROM inventories`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *InventoryRepository) Create(ctx context.Context, data inventory.Entity) (id string, err error) {
	query := `
        INSERT INTO inventories(store_id,product_id quantity, price)
        VALUES ($1, $2, $3,$4)
        RETURNING id`

	args := []interface{}{data.StoreID, data.ProductID, data.Quantity, data.Price}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *InventoryRepository) Get(ctx context.Context, id string) (dest inventory.Entity, err error) {
	query := `
        SELECT id, store_id,product_id quantity, quantity_min,quantity_max,price,price_sepcial,price_previous,is_available
        FROM inventories
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

func (s *InventoryRepository) Update(ctx context.Context, storeID string, data inventory.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, storeID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE inventories SET %s WHERE store_id=$%d", strings.Join(sets, ", "), len(args))
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

func (s *InventoryRepository) prepareArgs(data inventory.Entity) (sets []string, args []any) {

	if data.Quantity != nil {
		args = append(args, data.Quantity)
		sets = append(sets, fmt.Sprintf("quantity=$%d", len(args)))
	}
	if data.QuantityMin != nil {
		args = append(args, data.QuantityMin)
		sets = append(sets, fmt.Sprintf("quantity_min=$%d", len(args)))
	}
	if data.QuantityMax != nil {
		args = append(args, data.QuantityMax)
		sets = append(sets, fmt.Sprintf("quantity_max=$%d", len(args)))
	}
	if data.Price != nil {
		args = append(args, data.Price)
		sets = append(sets, fmt.Sprintf("price=$%d", len(args)))
	}
	if data.PriceSpecial != nil {
		args = append(args, data.PriceSpecial)
		sets = append(sets, fmt.Sprintf("price_special=$%d", len(args)))
	}
	if data.PricePrevious != nil {
		args = append(args, data.PricePrevious)
		sets = append(sets, fmt.Sprintf("price_previous=$%d", len(args)))
	}
	if data.IsAvailable != nil {
		args = append(args, data.IsAvailable)
		sets = append(sets, fmt.Sprintf("is_available=$%d", len(args)))
	}

	return
}

func (s *InventoryRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
        DELETE 
        FROM inventories
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
