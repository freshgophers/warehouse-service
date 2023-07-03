package memory

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"warehouse-service/internal/domain/inventory"
	"warehouse-service/pkg/storage"
)

type InventoryRepository struct {
	db map[string]inventory.Entity
	sync.RWMutex
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{
		db: make(map[string]inventory.Entity),
	}
}

func (r *InventoryRepository) Select(ctx context.Context) (dest []inventory.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]inventory.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *InventoryRepository) Create(ctx context.Context, data inventory.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil

}

func (r *InventoryRepository) Get(ctx context.Context, id string) (dest inventory.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = storage.ErrorNotFound
		return
	}

	return
}

func (r *InventoryRepository) Update(ctx context.Context, id string, data inventory.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	r.db[id] = data

	return
}

func (r *InventoryRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	delete(r.db, id)

	return
}

func (r *InventoryRepository) generateID() string {
	return uuid.New().String()
}
