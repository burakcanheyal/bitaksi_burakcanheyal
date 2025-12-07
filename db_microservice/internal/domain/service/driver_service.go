package service

import (
	"bitaksi_burakcanheyal/db_microservice/internal/application"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/dto"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/entity"
	"bitaksi_burakcanheyal/db_microservice/platform/mongo/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
// VALID HELPERS → TYPED ERRORS
// ──────────────────────────────────────────────
var allowedTaxiTypes = map[string]bool{
	"sarı":    true,
	"turkuaz": true,
	"siyah":   true,
}

func validateTaxiType(t string) error {
	if t == "" {
		return nil // opsiyonel
	}
	if !allowedTaxiTypes[t] {
		return application.Wrap("Taksi Tipi Hatası")
	}
	return nil
}

func validateCoordinates(lat, lon float64) error {
	if lat == 0 && lon == 0 {
		return nil // opsiyonel
	}
	if lat < -90 || lat > 90 {
		return application.Wrap("Koordinat Hatası: Lat")
	}
	if lon < -180 || lon > 180 {
		return application.Wrap("Koordinat Hatası : Lon")
	}
	return nil
}

// ──────────────────────────────────────────────
// CREATE DRIVER (VALIDATED)
// ──────────────────────────────────────────────
func (s *DriverService) CreateDriver(ctx context.Context, req dto.CreateDriverRequest) (string, error) {

	// TaxiType validation
	if err := validateTaxiType(req.TaxiType); err != nil {
		return "", err
	}

	// Coordinates validation
	if err := validateCoordinates(req.Lat, req.Lon); err != nil {
		return "", err
	}

	// Plate unique check
	existing, _ := s.repo.FindByPlate(ctx, req.Plate)
	if existing != nil {
		return "", application.Wrap("Kayıtlı Plaka") // “plate exists” → validation error
	}

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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := s.repo.Create(ctx, driver)
	if err != nil {
		return "", application.Wrap("ERR_INTERNAL")
	}

	return id, nil
}

// ──────────────────────────────────────────────
// UPDATE DRIVER
// ──────────────────────────────────────────────
func (s *DriverService) UpdateDriver(ctx context.Context, id string, req dto.UpdateDriverRequest) error {

	// ID validation
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		return application.Wrap("Hatalı ID")
	}

	// Existing driver fetch
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return application.Wrap("Sürücü Bulunamadı")
	}

	// Update optional fields
	if req.FirstName != "" {
		existing.FirstName = req.FirstName
	}
	if req.LastName != "" {
		existing.LastName = req.LastName
	}

	// TaxiType validation
	if err := validateTaxiType(req.TaxiType); err != nil {
		return err
	}
	if req.TaxiType != "" {
		existing.TaxiType = req.TaxiType
	}

	// Plate update
	if req.Plate != existing.Plate {
		other, _ := s.repo.FindByPlate(ctx, req.Plate)
		if other != nil && other.ID.Hex() != id {
			return application.Wrap("Kayıtlı Plaka") // “plate exists”
		}
		existing.Plate = req.Plate
	}

	if req.CarBrand != "" {
		existing.CarBrand = req.CarBrand
	}
	if req.CarModel != "" {
		existing.CarModel = req.CarModel
	}

	// Location update
	if req.Lat != 0 || req.Lon != 0 {
		if err := validateCoordinates(req.Lat, req.Lon); err != nil {
			return err
		}
		existing.Location.Coordinates = []float64{req.Lon, req.Lat}
	}

	existing.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, id, existing); err != nil {
		return application.Wrap("ERR_INTERNAL")
	}

	return nil
}

// ──────────────────────────────────────────────
// LIST
// ──────────────────────────────────────────────
func (s *DriverService) ListDrivers(ctx context.Context, page, pageSize int) ([]dto.DriverListResponse, error) {

	drivers, err := s.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, application.Wrap("ERR_INTERNAL")
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
			Lat:       d.Location.Coordinates[1],
			Lon:       d.Location.Coordinates[0],
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return res, nil
}

// ──────────────────────────────────────────────
// NEARBY DRIVERS
// ──────────────────────────────────────────────
func (s *DriverService) GetNearbyDrivers(ctx context.Context, req dto.NearbyDriverRequest) ([]dto.NearbyDriverResponse, error) {

	if err := validateCoordinates(req.Lat, req.Lon); err != nil {
		return nil, err
	}

	allDrivers, err := s.repo.ListAll(ctx)
	if err != nil {
		return nil, application.Wrap("ERR_INTERNAL")
	}

	var result []dto.NearbyDriverResponse

	for _, d := range allDrivers {

		// TaxiType filter (optional)
		if req.TaxiType != "" && d.TaxiType != req.TaxiType {
			continue
		}

		dist := haversine(req.Lat, req.Lon, d.Location.Coordinates[1], d.Location.Coordinates[0])

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

func degreesToRadians(d float64) float64 { return d * math.Pi / 180 }

func sortByDistance(list []dto.NearbyDriverResponse) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j].DistanceKm < list[i].DistanceKm {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
