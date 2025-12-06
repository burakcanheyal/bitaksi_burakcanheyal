package server

import (
	"bitaksi_burakcanheyal/driver_microservice/internal/application/handler"
	"bitaksi_burakcanheyal/driver_microservice/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGatewayRoutes(r *gin.Engine, h *handler.DriverGatewayHandler, log *handler.LogHandler) {

	g := r.Group("/drivers")

	g.Use(middleware.JwtAuthMiddleware())
	g.Use(middleware.RequestLogger())

	// CRUD + Nearby
	g.POST("", h.CreateDriver)
	g.PUT("/:id", h.UpdateDriver)
	g.GET("", h.ListDrivers)
	g.GET("/nearby", h.GetNearbyDrivers)

	internal := r.Group("/internal")
	internal.Use(middleware.JwtAuthMiddleware())
	internal.GET("/logs", log.GetLogs)

}
