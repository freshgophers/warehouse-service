package memory

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"warehouse-service/internal/domain/store"
	"warehouse-service/pkg/storage"
)

type StoreRepository struct {
	db map[string]store.Entity
	sync.RWMutex
}

func NewStoreRepository() *StoreRepository {
	return &StoreRepository{
		db: make(map[string]store.Entity),
	}
}

func (r *StoreRepository) Select(ctx context.Context) (dest []store.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]store.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *StoreRepository) Create(ctx context.Context, data store.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil

}

func (r *StoreRepository) Get(ctx context.Context, id string) (dest store.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = storage.ErrorNotFound
		return
	}

	return
}

func (r *StoreRepository) GetCityByID(ctx context.Context, id string) (dest store.City, err error) {
	return
}

func (r *StoreRepository) Update(ctx context.Context, id string, data store.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	r.db[id] = data

	return
}

func (r *StoreRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return storage.ErrorNotFound
	}
	delete(r.db, id)

	return
}

func (r *StoreRepository) generateID() string {
	return uuid.New().String()
}
