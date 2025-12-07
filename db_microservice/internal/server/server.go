package server

import (
	"bitaksi_burakcanheyal/db_microservice/internal/application/handler"
	"bitaksi_burakcanheyal/db_microservice/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, driverHandler *handler.DriverHandler) {
	api := r.Group("/drivers")
	api.Use(middleware.InternalApiKeyMiddleware()) // ← added
	api.Use(middleware.ErrorMapper())

	api.POST("", driverHandler.CreateDriver)
	api.PUT("/:id", driverHandler.UpdateDriver)
	api.GET("", driverHandler.ListDrivers)
	api.POST("/nearby", driverHandler.GetNearbyDrivers)
}
