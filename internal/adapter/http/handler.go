package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ic-x/blockchain-supply-service/internal/core/service"
)

// Handler struct holds a reference to the service.
type Handler struct {
	supplyService *service.SupplyService
}

// NewHandler creates a new Handler.
func NewHandler(svc *service.SupplyService) *Handler {
	return &Handler{supplyService: svc}
}

// GetTotalSupply is the HTTP handler for fetching the total supply.
// It only allows GET requests.
func (h *Handler) GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request to /getTotalSupply")

	// Receive data through the service
	supply, err := h.supplyService.GetTotalSupply()
	if err != nil {
		log.Printf("Error fetching total supply: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched supply: %+v", supply)

	// Set headers and return JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(supply); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

	log.Println("Response successfully returned")
}
