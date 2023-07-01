package warehouse

import (
	"context"
	"encoding/json"
	"warehouse-service/internal/domain/city"
	"warehouse-service/internal/domain/currency"
	"warehouse-service/internal/domain/delivery"
	"warehouse-service/internal/domain/schedule"
	"warehouse-service/internal/domain/store"
)

func (s *Service) ListStores(ctx context.Context) (res []store.Response, err error) {
	storeData, err := s.storeRepository.Select(ctx)
	if err != nil {
		return
	}
	res = store.ParseFromEntities(storeData)

	for i := 0; i < len(storeData); i++ {
		cityData, err := s.cityRepository.Get(ctx, storeData[i].CityID)
		if err != nil {
			return nil, err
		}
		res[i].City = city.ParseFromEntity(cityData)

		if cityData != nil {
			currencyData, err := s.currencyRepository.Get(ctx, cityData.CountryID)
			if err != nil {
				return nil, err
			}
			res[i].Currency = currency.ParseFromEntity(currencyData)

		}

		scheduleData, err := s.scheduleRepository.Get(ctx, storeData[i].ID)
		if err != nil {
			return nil, err
		}
		res[i].Schedule = schedule.ParseFromEntity(scheduleData)

		deliveryData, err := s.deliveryRepository.Get(ctx, storeData[i].ID)
		if err != nil {
			return nil, err
		}
		res[i].Delivery = delivery.ParseFromEntity(deliveryData)
	}

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
	}

	data.ID, err = s.storeRepository.Create(ctx, data)
	if err != nil {
		return
	}
	res = store.ParseFromEntity(data)

	return

}

func (s *Service) GetStore(ctx context.Context, id string) (res store.Response, err error) {
	storeData, err := s.storeRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = store.ParseFromEntity(storeData)

	cityData, err := s.cityRepository.Get(ctx, storeData.CityID)
	if err != nil {
		return
	}
	res.City = city.ParseFromEntity(cityData)

	if cityData != nil {
		currencyData, err := s.currencyRepository.Get(ctx, cityData.CountryID)
		if err != nil {
			return res, err
		}
		res.Currency = currency.ParseFromEntity(currencyData)
	}

	scheduleData, err := s.scheduleRepository.Get(ctx, storeData.ID)
	if err != nil {
		return
	}
	res.Schedule = schedule.ParseFromEntity(scheduleData)

	deliveryData, err := s.deliveryRepository.Get(ctx, storeData.ID)
	if err != nil {
		return
	}
	res.Delivery = delivery.ParseFromEntity(deliveryData)

	return
}

func (s *Service) UpdateStore(ctx context.Context, id string, req store.Request) (err error) {
	// Update store data by store ID
	storeData := store.Entity{
		Name:     &req.Name,
		Location: &req.Location,
		Rating:   &req.Rating,
		Address:  &req.Address,
		IsActive: &req.IsActive,
	}
	err = s.storeRepository.Update(ctx, id, storeData)
	if err != nil {
		return
	}
	cityData := city.Entity{
		Name:      &req.City.Name,
		GeoCenter: &req.City.GeoCenter,
	}

	err = s.cityRepository.Update(ctx, id, cityData)
	if err != nil {
		return
	}

	// get schedule data by store id
	scheduleData, err := s.scheduleRepository.Get(ctx, id)
	if err != nil {
		return
	}

	// marshal schedule periods to bytes
	periodsData, err := json.Marshal(req.Schedule.Periods)
	if err != nil {
		return
	}

	// check existing of schedule
	if scheduleData == nil {
		// if not exists create schedule data
		scheduleData := schedule.Entity{
			StoreID:  id,
			Periods:  periodsData,
			IsActive: true,
		}

		scheduleData.ID, err = s.scheduleRepository.Create(ctx, scheduleData)
		if err != nil {
			return
		}
		return
	}
	scheduleData.Periods = periodsData

	// if exists update schedule data
	err = s.scheduleRepository.Update(ctx, id, scheduleData)
	if err != nil {
		return
	}

	// get delivery data by store id
	deliveryData, err := s.deliveryRepository.Get(ctx, id)
	if err != nil {
		return
	}

	// marshal delivery periods to bytes
	areasData, err := json.Marshal(req.Delivery.Areas)
	if err != nil {
		return
	}

	periodsData, err = json.Marshal(req.Delivery.Periods)
	if err != nil {
		return
	}

	// check existing of schedule
	if deliveryData == nil {
		// if not exists create schedule data
		deliveryData := delivery.Entity{
			StoreID:  id,
			Periods:  periodsData,
			Areas:    areasData,
			IsActive: true,
		}

		deliveryData.ID, err = s.deliveryRepository.Create(ctx, deliveryData)
		if err != nil {
			return
		}
		return
	}
	deliveryData.Periods = periodsData
	deliveryData.Areas = areasData

	// if exists update schedule data
	err = s.deliveryRepository.Update(ctx, id, deliveryData)
	if err != nil {
		return
	}

	return

}

func (s *Service) DeleteStore(ctx context.Context, id string) (err error) {
	return s.storeRepository.Delete(ctx, id)
}
