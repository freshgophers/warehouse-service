package delivery

import "context"

type Repository interface {
	Create(ctx context.Context, data Entity) (dest string, err error)
	Get(ctx context.Context, storeID string) (dest *Entity, err error)
	Update(ctx context.Context, id string, data *Entity) (err error)
	Delete(ctx context.Context, id string) (err error)
}
