package supply_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ic-x/blockchain-supply-service/internal/core/domain"
)

type SupplyAPI struct {
	URL string
}

// External response format
type externalResponse struct {
	Supply []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"supply"`
}

// NewSupplyAPI creates a new instance of SupplyAPI.
func NewSupplyAPI(url string) *SupplyAPI {
	return &SupplyAPI{URL: url}
}

// GetTotalSupply fetches the total supply of NGL tokens from the external API.
func (api *SupplyAPI) GetTotalSupply() (domain.Supply, error) {
	log.Printf("Making request to external API: %s", api.URL)

	resp, err := http.Get(api.URL)
	if err != nil {
		return domain.Supply{}, fmt.Errorf("failed to fetch total supply: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Received response from external API with status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return domain.Supply{}, fmt.Errorf("received non-OK HTTP status: %d", resp.StatusCode)
	}

	var externalResp externalResponse
	if err := json.NewDecoder(resp.Body).Decode(&externalResp); err != nil {
		return domain.Supply{}, fmt.Errorf("failed to decode external response: %w", err)
	}

	log.Printf("External API response: %+v", externalResp)

	if len(externalResp.Supply) > 0 && externalResp.Supply[0].Denom == "aNGL" {
		return domain.Supply{Amount: externalResp.Supply[0].Amount}, nil
	}

	return domain.Supply{}, fmt.Errorf("aNGL token not found")
}
