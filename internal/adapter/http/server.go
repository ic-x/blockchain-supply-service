package http

import (
    "log"
    "net/http"

    "github.com/ic-x/blockchain-supply-service/internal/core/service"
)

// Server represents the HTTP server.
type Server struct {
    port          string
    supplyService *service.SupplyService
}

// NewServer creates a new HTTP server with the provided service and port.
func NewServer(port string, supplyService *service.SupplyService) *Server {
    return &Server{
        port:          port,
        supplyService: supplyService,
    }
}

// Start initializes and starts the HTTP server.
func (s *Server) Start() error {
    mux := http.NewServeMux()

    // Define routes using new syntax from Go 1.22
    handler := NewHandler(s.supplyService)
    mux.HandleFunc("GET /getTotalSupply", handler.GetTotalSupply)

    log.Printf("Starting server on port %s", s.port)
    return http.ListenAndServe(":"+s.port, mux)
}
