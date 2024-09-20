package port

import "github.com/ic-x/blockchain-supply-service/internal/core/domain"

// SupplyClient defines the interface (port) for fetching the total supply.
type SupplyClient interface {
	GetTotalSupply() (domain.Supply, error)
}
