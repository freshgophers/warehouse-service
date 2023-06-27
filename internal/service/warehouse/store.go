package warehouse

import (
	"context"
	"warehouse-service/internal/domain/store"
)

func (s *Service) ListStores(ctx context.Context) (res []store.Response, err error) {
	data, err := s.storeRepository.Select(ctx)
	if err != nil {
		return
	}
	res = store.ParseFromEntities(data)

	return
}

func (s *Service) AddStore(ctx context.Context, req store.Request) (res store.Response, err error) {
	data := store.Entity{
		MerchantID: req.MerchantID,
		CityID:     req.CityID,
		Name:       &req.Name,
		Address:    &req.Address,
		Location:   &req.Location,
		Rating:     &req.Rating,
		IsActive:   &req.IsActive,
	}

	data.ID, err = s.storeRepository.Create(ctx, data)
	if err != nil {
		return
	}
	res = store.ParseFromEntity(data)

	//data.MerchantID, err = s.storeRepository.Create(ctx, data)
	//if err != nil {
	//	return
	//}
	//res = store.ParseFromEntity(data)

	return

}

func (s *Service) GetStore(ctx context.Context, id string) (res store.Response, err error) {
	data, err := s.storeRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = store.ParseFromEntity(data)

	city, err := s.storeRepository.GetCityByID(ctx, data.CityID)
	if err != nil {
		return
	}
	res.City = city

	return
}

func (s *Service) UpdateStore(ctx context.Context, id string, req store.Request) (err error) {
	data := store.Entity{
		IsActive: &req.IsActive,
		Rating:   &req.Rating,
	}
	return s.storeRepository.Update(ctx, id, data)
}

func (s *Service) DeleteStore(ctx context.Context, id string) (err error) {
	return s.storeRepository.Delete(ctx, id)
}
