package repository

import (
	"warehouse-service/internal/domain/city"
	"warehouse-service/internal/domain/country"
	"warehouse-service/internal/domain/currency"
	"warehouse-service/internal/domain/delivery"
	"warehouse-service/internal/domain/inventory"
	"warehouse-service/internal/domain/schedule"
	"warehouse-service/internal/domain/store"
	"warehouse-service/internal/repository/memory"
	"warehouse-service/internal/repository/postgres"
	"warehouse-service/pkg/storage"
)

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	postgres *storage.Database

	Store store.Repository

	City city.Repository

	Schedule schedule.Repository

	Delivery delivery.Repository

	Country country.Repository

	Currency currency.Repository

	Inventory inventory.Repository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Create the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

// Close closes the repository and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (r *Repository) Close() {
	if r.postgres != nil {
		r.postgres.Client.Close()
	}
}

// WithMemoryStore applies a memory store to the Repository
func WithMemoryStore() Configuration {
	return func(s *Repository) (err error) {
		// Create the memory store, if we needed parameters, such as connection strings they could be inputted here

		s.Store = memory.NewStoreRepository()

		return
	}
}

// WithPostgresStore applies a postgres store to the Repository
func WithPostgresStore(schema, dataSourceName string) Configuration {
	return func(s *Repository) (err error) {
		// Create the postgres store, if we needed parameters, such as connection strings they could be inputted here
		s.postgres, err = storage.NewDatabase(schema, dataSourceName)
		if err != nil {
			return
		}

		err = s.postgres.Migrate()
		if err != nil {
			return
		}

		s.Store = postgres.NewStoreRepository(s.postgres.Client)

		s.City = postgres.NewCityRepository(s.postgres.Client)

		s.Schedule = postgres.NewScheduleRepository(s.postgres.Client)

		s.Delivery = postgres.NewDeliveryRepository(s.postgres.Client)

		s.Currency = postgres.NewCurrencyRepository(s.postgres.Client)

		s.Country = postgres.NewCountryRepository(s.postgres.Client)

		s.Inventory = postgres.NewInventoryRepository(s.postgres.Client)

		return
	}
}
