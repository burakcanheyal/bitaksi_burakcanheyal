package server

import (
	"bitaksi_burakcanheyal/db_microservice/internal/application/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, driverHandler *handler.DriverHandler) {
	driver := r.Group("/drivers")
	{
		driver.POST("", driverHandler.CreateDriver)
		driver.PUT("/:id", driverHandler.UpdateDriver)
		driver.GET("", driverHandler.ListDrivers)
		driver.GET("/nearby", driverHandler.GetNearbyDrivers)
	}
}
