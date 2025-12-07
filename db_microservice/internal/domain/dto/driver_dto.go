package dto

import "time"

//
// ──────────────────── CREATE ────────────────────
//

type CreateDriverRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Plate     string  `json:"plate" binding:"required"`
	TaxiType  string  `json:"taksiType" binding:"required"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat" binding:"required"` // INPUT Geo Lat
	Lon       float64 `json:"lon" binding:"required"` // INPUT Geo Lon
}

type CreateDriverResponse struct {
	ID string `json:"id"`
}

//
// ──────────────────── UPDATE ────────────────────
//

type UpdateDriverRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Plate     string  `json:"plate" binding:"required"`
	TaxiType  string  `json:"taksiType" binding:"required"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat" binding:"required"`
	Lon       float64 `json:"lon" binding:"required"`
}

//
// ──────────────────── LIST RESPONSE ────────────────────
//

// ⭐ OUTPUT: Client tarafına yine normal lat/lon dönüyoruz
type DriverListResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Plate     string    `json:"plate"`
	TaxiType  string    `json:"taksiType"`
	CarBrand  string    `json:"carBrand"`
	CarModel  string    `json:"carModel"`
	Lat       float64   `json:"lat"` // OUTPUT lat
	Lon       float64   `json:"lon"` // OUTPUT lon
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//
// ──────────────────── NEARBY ────────────────────
//

type NearbyDriverRequest struct {
	Lat      float64 `json:"lat" binding:"required"`
	Lon      float64 `json:"lon" binding:"required"`
	TaxiType string  `json:"taksiType" binding:"required"`
}

// OUTPUT DTO — burada lat/lon yok, sadece distance istenmiş
type NearbyDriverResponse struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Plate      string  `json:"plate"`
	TaxiType   string  `json:"taksiType"`
	DistanceKm float64 `json:"distanceKm"`
}
