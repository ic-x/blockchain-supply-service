package main

import (
	"log"

	"github.com/ic-x/blockchain-supply-service/internal/adapter/config"
	"github.com/ic-x/blockchain-supply-service/internal/adapter/external/supply_api"
	"github.com/ic-x/blockchain-supply-service/internal/adapter/http"
	"github.com/ic-x/blockchain-supply-service/internal/core/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Printf("Loaded config: %+v", cfg)

	// Initialize the external API client
	supplyAPIClient := supply_api.NewSupplyAPI(cfg.URL)

	// Initialize the service layer
	supplyService := service.NewSupplyService(supplyAPIClient)

	// Initialize and start the HTTP server
	server := http.NewServer(cfg.Port, supplyService)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
