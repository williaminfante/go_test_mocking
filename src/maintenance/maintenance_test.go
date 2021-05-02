package maintenance_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_mocking/src/maintenance"
)

func TestNeedsMaintenace(t *testing.T) {
	t.Run("when days is just 30", func(t *testing.T) {
		needsMaintenance, actualError := maintenance.NeedsMaintenace(30)
		assert.Equal(t, true, needsMaintenance)
		assert.Equal(t, nil, actualError)
	})
	t.Run("when days is above 30", func(t *testing.T) {
		needsMaintenance, actualError := maintenance.NeedsMaintenace(42)
		assert.Equal(t, true, needsMaintenance)
		assert.Equal(t, nil, actualError)
	})
	t.Run("when days is less than 30", func(t *testing.T) {
		needsMaintenance, actualError := maintenance.NeedsMaintenace(24)
		assert.Equal(t, false, needsMaintenance)
		assert.Equal(t, nil, actualError)
		zeroMaintenance, zeroError := maintenance.NeedsMaintenace(0)
		assert.Equal(t, false, zeroMaintenance)
		assert.Equal(t, nil, zeroError)
	})
	t.Run("when days is less than 0", func(t *testing.T) {
		expectedError := errors.New("cannot accept less than zero days")
		needsMaintenance, actualError := maintenance.NeedsMaintenace(-42)
		assert.Equal(t, false, needsMaintenance)
		assert.Equal(t, expectedError, actualError)
	})
}
