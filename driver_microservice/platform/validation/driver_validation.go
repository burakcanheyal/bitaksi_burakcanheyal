package validation

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddDriverRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Plate     string  `json:"plate" binding:"required"`
	TaxiType  string  `json:"taksiType" binding:"required"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat" binding:"required"`
	Lon       float64 `json:"lon" binding:"required"`
}
type UpdateDriverRequest struct {
	ID        string  `json:"id" binding:"required"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Plate     string  `json:"plate"`
	TaxiType  string  `json:"taksiType"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
}
type NearbyRequest struct {
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	TaxiType string  `json:"taksiType"`
}

func ValidateAddDriver(req AddDriverRequest) error {

	if len(req.FirstName) < 2 {
		return errors.New("firstName")
	}

	if len(req.LastName) < 2 {
		return errors.New("lastName")
	}

	if req.TaxiType != "sarı" && req.TaxiType != "turkuaz" && req.TaxiType != "siyah" {
		return errors.New("Taksi Tipi Hatası")
	}

	if req.Lat < -90 || req.Lat > 90 {
		return errors.New("Koordinat Hatası: Lat")
	}

	if req.Lon < -180 || req.Lon > 180 {
		return errors.New("Koordinat Hatası : Lon")
	}

	return nil
}
func ValidateListParams(page, size int) error {
	if page < 1 {
		return errors.New("page")
	}
	if size < 1 || size > 100 {
		return errors.New("pageSize")
	}
	return nil
}
func ValidateUpdateDriver(req UpdateDriverRequest) error {

	// ID zorunlu
	if req.ID == "" {
		return errors.New("ERR_MISSING_ID")
	}

	// MongoDB ObjectID format kontrolü
	if _, err := primitive.ObjectIDFromHex(req.ID); err != nil {
		return errors.New("Hatalı ID")
	}

	// Opsiyonel fakat boş olmaması daha doğru
	if req.FirstName != "" && len(req.FirstName) < 2 {
		return errors.New("firstName")
	}

	if req.LastName != "" && len(req.LastName) < 2 {
		return errors.New("lastName")
	}

	if req.TaxiType != "" &&
		req.TaxiType != "sarı" &&
		req.TaxiType != "turkuaz" &&
		req.TaxiType != "siyah" {

		return errors.New("Taksi Tipi Hatası")
	}

	// Eğer koordinatlar gönderilmişse valid olmalı
	if req.Lat != 0 {
		if req.Lat < -90 || req.Lat > 90 {
			return errors.New("Koordinat Hatası: Lat")
		}
	}

	if req.Lon != 0 {
		if req.Lon < -180 || req.Lon > 180 {
			return errors.New("Koordinat Hatası : Lon")
		}
	}

	return nil
}

func ValidateNearby(req NearbyRequest) error {

	// TaxiType kontrolü
	if req.TaxiType != "" &&
		req.TaxiType != "sarı" &&
		req.TaxiType != "turkuaz" &&
		req.TaxiType != "siyah" {

		return errors.New("Taksi Tipi Hatası")
	}

	// Lat gönderilmişse valid olmalı
	if req.Lat != 0 {
		if req.Lat < -90 || req.Lat > 90 {
			return errors.New("Koordinat Hatası: Lat")
		}
	}

	// Lon gönderilmişse valid olmalı
	if req.Lon != 0 {
		if req.Lon < -180 || req.Lon > 180 {
			return errors.New("Koordinat Hatası : Lon")
		}
	}

	return nil
}
