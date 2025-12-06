package main

import (
	"bitaksi_burakcanheyal/db_microservice/cmd"
	gateway "bitaksi_burakcanheyal/driver_microservice/cmd"
	"log"
	"net/http"
)

func main() {
	dbSrv := cmd.Db_Setup()
	testSrv := gateway.TestSetup()

	// Driver Service → 8080
	go func() {
		log.Println("[Driver-Service] running on :8080")
		if err := dbSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("db error: %v", err)
		}
	}()

	// Test Service → 9090
	go func() {
		log.Println("[Test-Service] running on :9090")
		if err := testSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("test error: %v", err)
		}
	}()

	select {} // keep alive
}
