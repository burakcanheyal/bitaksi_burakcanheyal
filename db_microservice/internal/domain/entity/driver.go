package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Plate     string             `json:"plate" bson:"plate"`
	TaxiType  string             `json:"taksiType" bson:"taksiType"`
	CarBrand  string             `json:"carBrand" bson:"carBrand"`
	CarModel  string             `json:"carModel" bson:"carModel"`
	Location  Location           `json:"location" bson:"location"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// ⭐ GEOJSON uyumlu Location modeli
type Location struct {
	Type        string    `json:"type" bson:"type"`               // "Point"
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [lon, lat]
}
