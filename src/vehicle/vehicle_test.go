package vehicle_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_mocking/src/vehicle"
)

func TestConductChecks(t *testing.T) {
	t.Run("no mocking testing", func(t *testing.T) {
		t.Run("maintenance", func(t *testing.T) {
			prescription, err := vehicle.ConductChecks("lightning", 45)
			assert.Equal(t,
				"lightning is up for maintenance",
				prescription)
			assert.Equal(t, nil, err)
		})
		t.Run("good", func(t *testing.T) {
			prescription, err := vehicle.ConductChecks("lightning", 23)
			assert.Equal(t,
				"lightning is all good",
				prescription)
			assert.Equal(t, nil, err)
		})
		t.Run("error", func(t *testing.T) {
			prescription, err := vehicle.ConductChecks("lightning", -45)
			assert.Equal(t, "", prescription)
			assert.Equal(t, errors.New("cannot accept less than zero days"), err)
		})
	})
}
