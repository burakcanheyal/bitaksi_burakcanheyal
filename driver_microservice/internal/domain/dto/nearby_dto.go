package dto

type NearbyRequest struct {
	Lat      float64 `json:"lat" binding:"required"`
	Lon      float64 `json:"lon" binding:"required"`
	TaxiType string  `json:"taksiType" binding:"required"`
}
