package memory

import (
	"context"
	"sync"
	"warehouse-service/pkg/storage"

	"github.com/google/uuid"

	"warehouse-service/internal/domain/product"
)

type ProductRepository struct {
	db map[string]product.Entity
	sync.RWMutex
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		db: make(map[string]product.Entity),
	}
}

func (r *ProductRepository) Select(ctx context.Context) (dest []product.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]product.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *ProductRepository) Create(ctx context.Context, data product.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil
}

func (r *ProductRepository) Get(ctx context.Context, id string) (dest product.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = storage.ErrorNotFound
		return
	}

	return
}

func (r *ProductRepository) Update(ctx context.Context, id string, data product.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	r.db[id] = data

	return
}

func (r *ProductRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	delete(r.db, id)

	return
}

func (r *ProductRepository) generateID() string {
	return uuid.New().String()
}
