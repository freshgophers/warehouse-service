package country

import "context"

type Repository interface {
	Select(ctx context.Context) (dest []Entity, err error)
	Create(ctx context.Context, data Entity) (dest string, err error)
	Get(ctx context.Context, id string) (dest Entity, err error)
	Update(ctx context.Context, id string, data Entity) (err error)
	Delete(ctx context.Context, id string) (err error)
}
