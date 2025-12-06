package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"bitaksi_burakcanheyal/db_microservice/internal/application/handler"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/service"
	"bitaksi_burakcanheyal/db_microservice/internal/server"
	pmongo "bitaksi_burakcanheyal/db_microservice/platform/mongo"
	"bitaksi_burakcanheyal/db_microservice/platform/mongo/repository"
	"github.com/gin-gonic/gin"
)

func Db_Setup() *http.Server {

	// ─────────────────────────────────────
	// 1) CONFIG
	// ─────────────────────────────────────
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "taxihub"
	}

	// ─────────────────────────────────────
	// 2) CONNECT MONGO
	// ─────────────────────────────────────
	client, err := pmongo.NewClient(mongoURI, dbName)
	if err != nil {
		log.Fatalf("mongo connection failed: %v", err)
	}

	db := client.Database(dbName)
	driverCollection := db.Collection("drivers")
	log.Println("[Mongo] Connected successfully.")

	// ─────────────────────────────────────
	// 3) REPOSITORY → SERVICE → HANDLER
	// ─────────────────────────────────────
	driverRepo := repository.NewDriverRepository(driverCollection)
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(driverService)

	// ─────────────────────────────────────
	// 4) ROUTER INIT
	// ─────────────────────────────────────
	r := gin.Default()
	server.RegisterRoutes(r, driverHandler)

	// ─────────────────────────────────────
	// 5) SERVER INIT (no async here!)
	// ─────────────────────────────────────
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv
}
