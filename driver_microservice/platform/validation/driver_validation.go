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
		return errors.New("firstName must be at least 2 characters")
	}

	if len(req.LastName) < 2 {
		return errors.New("lastName must be at least 2 characters")
	}

	if req.TaxiType != "yellow" && req.TaxiType != "turquoise" && req.TaxiType != "black" {
		return errors.New("invalid taksiType (allowed: yellow, turquoise, black)")
	}

	if req.Lat < -90 || req.Lat > 90 {
		return errors.New("lat must be between -90 and 90")
	}

	if req.Lon < -180 || req.Lon > 180 {
		return errors.New("lon must be between -180 and 180")
	}

	return nil
}
func ValidateListParams(page, size int) error {
	if page < 1 {
		return errors.New("page must be >= 1")
	}
	if size < 1 || size > 100 {
		return errors.New("pageSize must be between 1 and 100")
	}
	return nil
}
func ValidateUpdateDriver(req UpdateDriverRequest) error {

	// ID zorunlu
	if req.ID == "" {
		return errors.New("id is required")
	}

	// MongoDB ObjectID format kontrolü
	if _, err := primitive.ObjectIDFromHex(req.ID); err != nil {
		return errors.New("id must be valid MongoDB ObjectID")
	}

	// Opsiyonel fakat boş olmaması daha doğru
	if req.FirstName != "" && len(req.FirstName) < 2 {
		return errors.New("firstName must be at least 2 characters")
	}

	if req.LastName != "" && len(req.LastName) < 2 {
		return errors.New("lastName must be at least 2 characters")
	}

	if req.TaxiType != "" &&
		req.TaxiType != "yellow" &&
		req.TaxiType != "turquoise" &&
		req.TaxiType != "black" {

		return errors.New("invalid taksiType (allowed: yellow, turquoise, black)")
	}

	// Eğer koordinatlar gönderilmişse valid olmalı
	if req.Lat != 0 {
		if req.Lat < -90 || req.Lat > 90 {
			return errors.New("lat must be between -90 and 90")
		}
	}

	if req.Lon != 0 {
		if req.Lon < -180 || req.Lon > 180 {
			return errors.New("lon must be between -180 and 180")
		}
	}

	return nil
}

func ValidateNearby(req NearbyRequest) error {

	// TaxiType kontrolü
	if req.TaxiType != "" &&
		req.TaxiType != "yellow" &&
		req.TaxiType != "turquoise" &&
		req.TaxiType != "black" {

		return errors.New("invalid taksiType")
	}

	// Lat gönderilmişse valid olmalı
	if req.Lat != 0 {
		if req.Lat < -90 || req.Lat > 90 {
			return errors.New("lat must be between -90 and 90")
		}
	}

	// Lon gönderilmişse valid olmalı
	if req.Lon != 0 {
		if req.Lon < -180 || req.Lon > 180 {
			return errors.New("lon must be between -180 and 180")
		}
	}

	return nil
}
