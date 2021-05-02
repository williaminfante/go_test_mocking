package vehicle

import (
	"fmt"

	"github.com/williaminfante/go_test_mocking/maintenance"
)

func ConductChecks(name string, days_used int) (string, error) {
	decision, errMaintenance := maintenance.NeedsMaintenace(days_used)
	if errMaintenance != nil {
		return "", errMaintenance
	}

	if decision {
		return fmt.Sprintf("%v is up for maintenance", name), nil
	} else {
		return fmt.Sprintf("%v is all good", name), nil
	}
}
