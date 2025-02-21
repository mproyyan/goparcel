package app

import (
	"context"
	"errors"
	"testing"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/couriers/domain"
	"github.com/mproyyan/goparcel/internal/couriers/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestGetAvailableCouriers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	courierRepo := mock.NewMockCourierRepository(ctrl)
	courierService := NewCourierService(courierRepo)
	ctx := context.Background()

	tests := []struct {
		name          string
		locationID    string
		setupMock     func(locationID string)
		expectedError error
	}{
		// Get available couriers success
		{
			name:       "Get available couriers success",
			locationID: primitive.NewObjectID().Hex(),
			setupMock: func(locationID string) {
				id, _ := primitive.ObjectIDFromHex(locationID)
				courierRepo.EXPECT().AvailableCouriers(ctx, id).Return([]domain.Courier{}, nil)
			},
		},
		// Get available couriers failed
		{
			name:       "Get available couriers failed",
			locationID: primitive.NewObjectID().Hex(),
			setupMock: func(locationID string) {
				id, _ := primitive.ObjectIDFromHex(locationID)
				courierRepo.EXPECT().AvailableCouriers(ctx, id).Return([]domain.Courier{}, errors.New("failed to get couriers"))
			},
			expectedError: cuserr.Decorate(errors.New("failed to get couriers"), "failed to get available couriers"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(test.locationID)
			_, err := courierService.AvailableCouriers(ctx, test.locationID)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
