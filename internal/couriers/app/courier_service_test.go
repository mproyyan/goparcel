package app

import (
	"context"
	"errors"
	"testing"

	"github.com/mproyyan/goparcel/internal/couriers/domain"
	"github.com/mproyyan/goparcel/internal/couriers/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewCourierService(t *testing.T) {
	tests := []struct {
		name              string
		courierRepository domain.CourierRepository
		wantNil           bool
	}{
		{
			name: "Success - Create service with valid repository",
			courierRepository: func() domain.CourierRepository {
				ctrl := gomock.NewController(t)
				return mocks.NewMockCourierRepository(ctrl)
			}(),
			wantNil: false,
		},
		{
			name:              "Success - Create service with nil repository",
			courierRepository: nil,
			wantNil:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			service := NewCourierService(tt.courierRepository)

			// Assert
			if tt.wantNil {
				assert.Nil(t, service.courierRepository)
			} else {
				assert.Equal(t, tt.courierRepository, service.courierRepository)
			}

			// Verify service is not zero value
			assert.NotNil(t, service)
		})
	}
}

func TestCourierService_AvailableCouriers(t *testing.T) {
	tests := []struct {
		name           string
		locationID     string
		setupMock      func(*mocks.MockCourierRepository)
		expectedResult []domain.Courier
		expectedError  string
		wantError      bool
	}{
		{
			name:       "Success - Single available courier",
			locationID: "location-123",
			setupMock: func(mockRepo *mocks.MockCourierRepository) {
				expectedCouriers := []domain.Courier{
					{
						ID:         "courier-1",
						UserID:     "user-1",
						Name:       "John Doe",
						Email:      "john@example.com",
						Status:     domain.Available,
						LocationID: "location-123",
					},
				}

				mockRepo.EXPECT().AvailableCouriers(gomock.Any(), "location-123").Return(expectedCouriers, nil)
			},
			expectedResult: []domain.Courier{
				{
					ID:         "courier-1",
					UserID:     "user-1",
					Name:       "John Doe",
					Email:      "john@example.com",
					Status:     domain.Available,
					LocationID: "location-123",
				},
			},
			expectedError: "",
			wantError:     false,
		},
		{
			name:       "Success - Multiple available couriers",
			locationID: "location-456",
			setupMock: func(mockRepo *mocks.MockCourierRepository) {
				expectedCouriers := []domain.Courier{
					{
						ID:         "courier-1",
						UserID:     "user-1",
						Name:       "John Doe",
						Email:      "john@example.com",
						Status:     domain.Available,
						LocationID: "location-456",
					},
					{
						ID:         "courier-2",
						UserID:     "user-2",
						Name:       "Jane Smith",
						Email:      "jane@example.com",
						Status:     domain.Available,
						LocationID: "location-456",
					},
					{
						ID:         "courier-3",
						UserID:     "user-3",
						Name:       "Bob Wilson",
						Email:      "bob@example.com",
						Status:     domain.Available,
						LocationID: "location-456",
					},
				}
				mockRepo.EXPECT().AvailableCouriers(gomock.Any(), "location-456").Return(expectedCouriers, nil)
			},
			expectedResult: []domain.Courier{
				{
					ID:         "courier-1",
					UserID:     "user-1",
					Name:       "John Doe",
					Email:      "john@example.com",
					Status:     domain.Available,
					LocationID: "location-456",
				},
				{
					ID:         "courier-2",
					UserID:     "user-2",
					Name:       "Jane Smith",
					Email:      "jane@example.com",
					Status:     domain.Available,
					LocationID: "location-456",
				},
				{
					ID:         "courier-3",
					UserID:     "user-3",
					Name:       "Bob Wilson",
					Email:      "bob@example.com",
					Status:     domain.Available,
					LocationID: "location-456",
				},
			},
			expectedError: "",
			wantError:     false,
		},
		{
			name:       "Success - Empty result (no available couriers)",
			locationID: "location-789",
			setupMock: func(mockRepo *mocks.MockCourierRepository) {
				mockRepo.EXPECT().AvailableCouriers(gomock.Any(), "location-789").Return([]domain.Courier{}, nil)
			},
			expectedResult: []domain.Courier{},
			expectedError:  "",
			wantError:      false,
		},
		{
			name:       "Success - Nil result (no available couriers)",
			locationID: "location-999",
			setupMock: func(mockRepo *mocks.MockCourierRepository) {
				mockRepo.EXPECT().AvailableCouriers(gomock.Any(), "location-999").Return(nil, nil)
			},
			expectedResult: nil,
			expectedError:  "",
			wantError:      false,
		},
		{
			name:       "Error - Repository returns database error",
			locationID: "location-123",
			setupMock: func(mockRepo *mocks.MockCourierRepository) {
				mockRepo.EXPECT().AvailableCouriers(gomock.Any(), "location-123").Return(nil, errors.New("database connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get available couriers, cause: database connection failed",
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockCourierRepository(ctrl)
			courierService := CourierService{
				courierRepository: mockRepo,
			}

			// Setup mock expectations
			tt.setupMock(mockRepo)

			// Act
			ctx := context.Background()
			result, err := courierService.AvailableCouriers(ctx, tt.locationID)

			// Assert
			if tt.wantError {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.Equal(t, tt.expectedResult, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
