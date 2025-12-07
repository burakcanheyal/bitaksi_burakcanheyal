package cmd

import (
	"bitaksi_burakcanheyal/driver_microservice/internal/application/handler"
	"bitaksi_burakcanheyal/driver_microservice/internal/domain/model"
	"bitaksi_burakcanheyal/driver_microservice/internal/server"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GatewaySetup() *http.Server {

	driverClient := model.NewDriverClient("http://localhost:8080")

	driverHandler := handler.NewDriverGatewayHandler(driverClient)
	logHandler := handler.NewLogHandler()

	r := gin.Default()

	server.RegisterGatewayRoutes(r, driverHandler, logHandler)

	srv := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv
}
