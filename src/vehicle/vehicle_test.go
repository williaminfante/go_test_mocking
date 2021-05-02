package vehicle_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_mocking/mocks"
	"github.com/williaminfante/go_test_mocking/src/vehicle"
)

func TestConductChecks(t *testing.T) {
	testM := new(mocks.MockedMaintenanceWrapper)
	maintenanceRepository := vehicle.MaintenanceWrapper{Maintenance: testM}
	t.Run("no mocking testing", func(t *testing.T) {
		t.Run("maintenance", func(t *testing.T) {
			testM.On("NeedsMaintenance", 45).Return(true).Once()
			prescription, err := maintenanceRepository.ConductChecks("lightning", 45)
			testM.AssertExpectations(t)
			assert.Equal(t,
				"lightning is up for maintenance",
				prescription)
			assert.Equal(t, nil, err)
		})
		t.Run("good", func(t *testing.T) {
			testM.On("NeedsMaintenance", 23).Return(false).Once()
			prescription, err := maintenanceRepository.ConductChecks("lightning", 23)
			testM.AssertExpectations(t)
			assert.Equal(t,
				"lightning is all good",
				prescription)
			assert.Equal(t, nil, err)
		})
		t.Run("error", func(t *testing.T) {
			testM.On("NeedsMaintenance", -45).Return(false, errors.New("Sample")).Once()
			prescription, err := maintenanceRepository.ConductChecks("lightning", -45)
			testM.AssertExpectations(t)
			assert.Equal(t, "", prescription)
			assert.Equal(t, errors.New("Sample"), err)
		})
	})
}
