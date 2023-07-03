package warehouse

import (
	"warehouse-service/internal/domain/city"
	"warehouse-service/internal/domain/country"
	"warehouse-service/internal/domain/currency"
	"warehouse-service/internal/domain/delivery"
	"warehouse-service/internal/domain/inventory"
	"warehouse-service/internal/domain/schedule"
	"warehouse-service/internal/domain/store"
)

// Configuration is an alias for a function that will take in a pointer to a Service and modify it
type Configuration func(s *Service) error

// Service is an implementation of the Service
type Service struct {
	storeCache      store.Cache
	storeRepository store.Repository

	cityRepository city.Repository

	scheduleRepository schedule.Repository

	deliveryRepository delivery.Repository

	countryRepository country.Repository

	currencyRepository currency.Repository

	inventoryRepository inventory.Repository
	inventoryCache      inventory.Cache
}

// New takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Service, err error) {
	// Create the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

// WithStoreRepository applies a given category repository to the Service
func WithStoreRepository(storeRepository store.Repository) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.storeRepository = storeRepository
		return nil
	}
}

func WithInventoryRepository(inventoryRepository inventory.Repository) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.inventoryRepository = inventoryRepository
		return nil
	}
}

// WithStoreCache applies a given product cache to the Service
func WithStoreCache(storeCache store.Cache) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.storeCache = storeCache
		return nil
	}
}

func WithInventoryCache(inventoryCache inventory.Cache) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.inventoryCache = inventoryCache
		return nil
	}
}
