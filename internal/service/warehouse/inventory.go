package warehouse

import (
	"context"
	"warehouse-service/internal/domain/inventory"
)

func (s *Service) ListInventory(ctx context.Context) (res []inventory.Response, err error) {
	inventorData, err := s.inventoryRepository.Select(ctx)
	if err != nil {
		return
	}
	res = inventory.ParseFromEntities(inventorData)

	return
}

func (s *Service) AddInventory(ctx context.Context, req inventory.Request) (res inventory.Response, err error) {
	data := inventory.Entity{
		StoreID:   req.StoreID,
		ProductID: req.ProductID,
		Quantity:  &req.Quantity,
		Price:     &req.Price,
	}

	data.ID, err = s.inventoryRepository.Create(ctx, data)
	if err != nil {
		return
	}
	res = inventory.ParseFromEntity(data)

	return
}

func (s *Service) GetInventory(ctx context.Context, id string) (res inventory.Response, err error) {
	inventoryData, err := s.inventoryRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = inventory.ParseFromEntity(inventoryData)
	return
}

func (s *Service) UpdateInventory(ctx context.Context, id string, req inventory.Request) (err error) {
	data := inventory.Entity{
		Quantity:      &req.Quantity,
		QuantityMin:   req.QuantityMin,
		QuantityMax:   req.QuantityMax,
		Price:         &req.Price,
		PriceSpecial:  req.PriceSpecial,
		PricePrevious: req.PricePrevious,
	}
	return s.inventoryRepository.Update(ctx, id, data)
}

func (s *Service) DeleteInventory(ctx context.Context, id string) (err error) {
	return s.inventoryRepository.Delete(ctx, id)
}
