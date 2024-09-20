package service

import (
	"github.com/ic-x/blockchain-supply-service/internal/core/domain"
	"github.com/ic-x/blockchain-supply-service/internal/core/port"
)

// SupplyService handles business logic related to NGL token supply.
type SupplyService struct {
	client port.SupplyClient
}

// NewSupplyService creates a new SupplyService.
func NewSupplyService(client port.SupplyClient) *SupplyService {
	return &SupplyService{client: client}
}

// GetTotalSupply returns the total supply of NGL tokens.
func (s *SupplyService) GetTotalSupply() (domain.Supply, error) {
	return s.client.GetTotalSupply()
}
