package app

import (
	"context"
	"errors"
	"testing"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
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
		{
			name: "Fail because warehouse has warehouse ID",
			input: domain.Location{
				Name:        "Warehouse A",
				Type:        domain.Warehouse,
				WarehouseID: "wh-123",
				Address: domain.Address{
					ZipCode: "12345",
				},
			},
			mockRegion:    func() {},
			mockRepo:      func() {},
			expectedError: status.Error(codes.InvalidArgument, "warehouse type location cannot contain another warehouse id"),
		},
		{
			name: "Fail due to invalid location type",
			input: domain.Location{
				Name: "Invalid Location",
				Address: domain.Address{
					ZipCode: "12345",
				},
			},
			mockRegion:    func() {},
			mockRepo:      func() {},
			expectedError: status.Error(codes.InvalidArgument, "invalid location type, must be depot or warehouse"),
		},
		{
			name: "Fail when region service returns error",
			input: domain.Location{
				Name: "Warehouse C",
				Type: domain.Warehouse,
				Address: domain.Address{
					ZipCode: "99999",
				},
			},
			mockRegion: func() {
				mockRegionService.EXPECT().
					GetRegion(ctx, "99999").
					Return(nil, errors.New("region not found"))
			},
			mockRepo:      func() {},
			expectedError: cuserr.Decorate(errors.New("region not found"), "RegionService failed"),
		},
		{
			name: "Fail when repository CreateLocation fails",
			input: domain.Location{
				Name:        "Depot D",
				Type:        domain.Depot,
				WarehouseID: "wh-123",
				Address: domain.Address{
					ZipCode: "12345",
				},
			},
			mockRegion: func() {
				mockRegionService.EXPECT().
					GetRegion(ctx, "12345").
					Return(&domain.Region{
						Province:    "Province A",
						City:        "City A",
						District:    "District A",
						Subdistrict: "Subdistrict A",
					}, nil)
			},
			mockRepo: func() {
				mockLocationRepository.EXPECT().
					CreateLocation(ctx, gomock.Any()).
					Return("", errors.New("db error"))
			},
			expectedError: cuserr.Decorate(errors.New("db error"), "repository CreateLocation failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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

func TestGetLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepository := mock.NewMockLocationRepository(ctrl)
	locationService := NewLocationService(nil, locationRepository)
	ctx := context.Background()

	tests := []struct {
		name           string
		locationID     string
		mockRepo       func()
		expectedResult *domain.Location
		expectedError  error
	}{
		{
			name:       "Location found",
			locationID: "loc123",
			mockRepo: func() {
				locationRepository.EXPECT().
					FindLocation(ctx, gomock.Any()).
					Return(&domain.Location{ID: "loc123", Name: "X"}, nil)
			},
			expectedResult: &domain.Location{ID: "loc123", Name: "X"},
		},
		{
			name:       "Location not found",
			locationID: "loc123",
			mockRepo: func() {
				locationRepository.EXPECT().
					FindLocation(ctx, gomock.Any()).
					Return(nil, errors.New("location not found"))
			},
			expectedError: cuserr.Decorate(errors.New("location not found"), "repository FindLocation failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockRepo()
			result, err := locationService.GetLocation(ctx, test.locationID)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expectedResult, result)
			}
		})
	}
}

func TestGetRegion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegionService := mock.NewMockRegionService(ctrl)
	locationService := NewLocationService(mockRegionService, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		zipcode       string
		setupMock     func(zipcode string)
		expectedError error
	}{
		// Get region success
		{
			name:    "Get region success",
			zipcode: "11111",
			setupMock: func(zipcode string) {
				mockRegionService.EXPECT().GetRegion(ctx, zipcode).Return(&domain.Region{ZipCode: zipcode}, nil)
			},
		},
		// Get region failed
		{
			name:    "Get region failed",
			zipcode: "55555",
			setupMock: func(zipcode string) {
				mockRegionService.EXPECT().GetRegion(ctx, zipcode).Return(nil, errors.New("region not found"))
			},
			expectedError: cuserr.Decorate(errors.New("region not found"), "RegionService failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(test.zipcode)
			region, err := locationService.GetRegion(ctx, test.zipcode)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.zipcode, region.ZipCode)
			}
		})
	}
}
