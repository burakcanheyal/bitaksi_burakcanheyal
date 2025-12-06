package service

import (
	"bitaksi_burakcanheyal/db_microservice/internal/domain/dto"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/entity"
	"bitaksi_burakcanheyal/db_microservice/platform/mongo/repository"
	"context"
	"errors"
	"math"
	"time"
)

type DriverService struct {
	repo *repository.DriverRepository
}

func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{repo: repo}
}

// ──────────────────────────────────────────────
// CREATE DRIVER
// ──────────────────────────────────────────────
func (s *DriverService) CreateDriver(ctx context.Context, req dto.CreateDriverRequest) (string, error) {

	if req.FirstName == "" || req.LastName == "" || req.Plate == "" || req.TaxiType == "" {
		return "", errors.New("missing required fields")
	}

	driver := &entity.Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		Location: entity.Location{
			Type:        "Point",                     // ⭐ GEO JSON TYPE
			Coordinates: []float64{req.Lon, req.Lat}, // ⭐ [lon, lat] formatı
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.Create(ctx, driver)
}

// ──────────────────────────────────────────────
// UPDATE DRIVER
// ──────────────────────────────────────────────
func (s *DriverService) UpdateDriver(ctx context.Context, id string, req dto.UpdateDriverRequest) error {

	driver := &entity.Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		Location: entity.Location{
			Type:        "Point",
			Coordinates: []float64{req.Lon, req.Lat},
		},
		UpdatedAt: time.Now(),
	}

	return s.repo.Update(ctx, id, driver)
}

// ──────────────────────────────────────────────
// LIST DRIVERS
// ──────────────────────────────────────────────
func (s *DriverService) ListDrivers(ctx context.Context, page, pageSize int) ([]dto.DriverListResponse, error) {

	drivers, err := s.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	var res []dto.DriverListResponse
	for _, d := range drivers {
		res = append(res, dto.DriverListResponse{
			ID:        d.ID.Hex(),
			FirstName: d.FirstName,
			LastName:  d.LastName,
			Plate:     d.Plate,
			TaxiType:  d.TaxiType,
			CarBrand:  d.CarBrand,
			CarModel:  d.CarModel,
			Lat:       d.Location.Coordinates[1], // ⭐ lat index 1
			Lon:       d.Location.Coordinates[0], // ⭐ lon index 0
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return res, nil
}

// ──────────────────────────────────────────────
// NEARBY DRIVERS (Haversine)
// ──────────────────────────────────────────────
func (s *DriverService) GetNearbyDrivers(ctx context.Context, req dto.NearbyDriverRequest) ([]dto.NearbyDriverResponse, error) {

	allDrivers, err := s.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.NearbyDriverResponse

	for _, d := range allDrivers {
		if d.TaxiType != req.TaxiType {
			continue
		}

		driverLat := d.Location.Coordinates[1]
		driverLon := d.Location.Coordinates[0]

		dist := haversine(req.Lat, req.Lon, driverLat, driverLon)

		if dist <= 6.0 {
			result = append(result, dto.NearbyDriverResponse{
				FirstName:  d.FirstName,
				LastName:   d.LastName,
				Plate:      d.Plate,
				TaxiType:   d.TaxiType,
				DistanceKm: dist,
			})
		}
	}

	sortByDistance(result)
	return result, nil
}

// ──────────────────────────────────────────────
// UTILITIES
// ──────────────────────────────────────────────
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371

	dLat := degreesToRadians(lat2 - lat1)
	dLon := degreesToRadians(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(degreesToRadians(lat1))*math.Cos(degreesToRadians(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func sortByDistance(list []dto.NearbyDriverResponse) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j].DistanceKm < list[i].DistanceKm {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
