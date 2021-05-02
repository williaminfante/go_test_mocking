package vehicle_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_mocking/mocks"
	"github.com/williaminfante/go_test_mocking/src/vehicle"
)

func TestConductChecks(t *testing.T) {
	testWrapper := new(mocks.MockedMaintenanceWrapper)
	maintenanceRepository := vehicle.MaintenanceWrapper{Maintenance: testWrapper}
	t.Run("maintenance", func(t *testing.T) {
		testDays := 45
		testWrapper.On("NeedsMaintenance", testDays).Return(true).Once()
		prescription, err := maintenanceRepository.ConductChecks("lightning", testDays)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, "lightning is up for maintenance", prescription)
		assert.Equal(t, nil, err)
	})
	t.Run("good", func(t *testing.T) {
		testDays := 23
		testWrapper.On("NeedsMaintenance", testDays).Return(false).Once()
		prescription, err := maintenanceRepository.ConductChecks("lightning", testDays)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, "lightning is all good", prescription)
		assert.Equal(t, nil, err)
	})
	t.Run("error", func(t *testing.T) {
		testDays := -45
		expectedErr := errors.New("Sample")
		testWrapper.On("NeedsMaintenance", testDays).Return(false, expectedErr).Once()
		prescription, err := maintenanceRepository.ConductChecks("lightning", testDays)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, "", prescription)
		assert.Equal(t, expectedErr, err)
	})
}
