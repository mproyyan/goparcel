package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mproyyan/goparcel/internal/locations/domain"
)

type RegionService struct {
	client *http.Client
	apiKey string
}

func NewRegionService(client *http.Client, apiKey string) *RegionService {
	return &RegionService{
		client: client,
		apiKey: apiKey,
	}
}

func (r *RegionService) GetRegion(ctx context.Context, zipcode string) (*domain.Region, error) {
	// Create request
	url := fmt.Sprintf("https://rajaongkir.komerce.id/api/v1/destination/domestic-destination?search=%s", zipcode)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add api key
	req.Header.Set("key", r.apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check http statis code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parsing JSON response
	var result struct {
		Data []RegionModel `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// If no data found
	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no region found for zipcode: %s", zipcode)
	}

	// If only one data found then include subdistrict
	if len(result.Data) == 1 {
		domainRegion := regionModelToDomain(result.Data[0])
		return &domainRegion, nil
	}

	// If found more than one data then remove subdistrict
	region := result.Data[0]
	region.Subdistrict = ""

	domainRegion := regionModelToDomain(region)
	return &domainRegion, nil
}

// Models
type RegionModel struct {
	ZipCode     string `json:"zip_code"`
	Province    string `json:"province_name"`
	City        string `json:"city_name"`
	District    string `json:"district_name"`
	Subdistrict string `json:"subdistrict_name"`
}

func regionModelToDomain(regionModel RegionModel) domain.Region {
	return domain.Region{
		ZipCode:     regionModel.ZipCode,
		Province:    regionModel.Province,
		City:        regionModel.City,
		District:    regionModel.District,
		Subdistrict: regionModel.Subdistrict,
	}
}
