package app

import (
	"context"
	"testing"

	"github.com/mproyyan/goparcel/internal/locations/domain"
	"github.com/mproyyan/goparcel/internal/locations/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegionService := mock.NewMockRegionService(ctrl)
	mockLocationRepository := mock.NewMockLocationRepository(ctrl)

	locationService := NewLocationService(mockRegionService, mockLocationRepository)
	ctx := context.Background()

	// Test table
	tests := []struct {
		name          string
		input         domain.Location
		mockRegion    func()
		mockRepo      func()
		expectedError error
	}{
		{
			name: "Success create depot",
			input: domain.Location{
				Name:        "X",
				Type:        domain.Depot,
				WarehouseID: "123",
				Address: domain.Address{
					ZipCode: "11111",
				},
			},
			mockRegion: func() {
				mockRegionService.EXPECT().
					GetRegion(ctx, "11111").
					Return(&domain.Region{
						ZipCode:     "11111",
						Province:    "Province A",
						City:        "City A",
						District:    "District A",
						Subdistrict: "Subdistrict A",
					}, nil)
			},
			mockRepo: func() {
				mockLocationRepository.EXPECT().
					CreateLocation(ctx, domain.Location{
						Name:        "X",
						Type:        domain.Depot,
						WarehouseID: "123",
						Address: domain.Address{
							ZipCode:     "11111",
							Province:    "Province A",
							City:        "City A",
							District:    "District A",
							Subdistrict: "Subdistrict A",
						},
					}).
					Return("12345", nil)
			},
			expectedError: nil,
		},
		{
			name: "Success create warehouse",
			input: domain.Location{
				Name: "X",
				Type: domain.Warehouse,
				Address: domain.Address{
					ZipCode: "11111",
				},
			},
			mockRegion: func() {
				mockRegionService.EXPECT().
					GetRegion(ctx, "11111").
					Return(&domain.Region{
						ZipCode:     "11111",
						Province:    "Province A",
						City:        "City A",
						District:    "District A",
						Subdistrict: "Subdistrict A",
					}, nil)
			},
			mockRepo: func() {
				mockLocationRepository.EXPECT().
					CreateLocation(ctx, domain.Location{
						Name: "X",
						Type: domain.Warehouse,
						Address: domain.Address{
							ZipCode:     "11111",
							Province:    "Province A",
							City:        "City A",
							District:    "District A",
							Subdistrict: "Subdistrict A",
						},
					}).
					Return("12345", nil)
			},
			expectedError: nil,
		},
		{
			name: "Fail because depot has no warehouse ID",
			input: domain.Location{
				Name:        "Depot B",
				Type:        domain.Depot,
				WarehouseID: "",
				Address: domain.Address{
					ZipCode: "12345",
				},
			},
			mockRegion:    func() {},
			mockRepo:      func() {},
			expectedError: status.Error(codes.InvalidArgument, "depot type location must contain warehouse id"),
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			test.mockRegion()
			test.mockRepo()

			err := locationService.CreateLocation(ctx, test.input)

			if test.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			}
		})
	}
}

func newTestLocationData(locationType string, region *domain.Region, warehouseID string) domain.Location {
	location := domain.Location{
		Name:        "X",
		Type:        domain.LocationTypeFromString(locationType),
		WarehouseID: warehouseID,
	}

	location.Address.Province = region.Province
	location.Address.City = region.City
	location.Address.District = region.District
	location.Address.Subdistrict = region.Subdistrict
	location.Address.ZipCode = region.ZipCode

	return location
}

func newTestRegionData() *domain.Region {
	return &domain.Region{
		ZipCode:     "11111",
		Province:    "Province A",
		City:        "City A",
		District:    "District A",
		Subdistrict: "Subdistrict A",
	}
}
