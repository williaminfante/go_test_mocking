package vehicle

import (
	"fmt"

	"github.com/williaminfante/go_test_mocking/src/maintenance"
)

func ConductChecks(name string, days_used int) (string, error) {
	specialist := maintenance.Specialist{}
	decision, errMaintenance := specialist.NeedsMaintenance(days_used)
	if errMaintenance != nil {
		return "", errMaintenance
	}
	if decision {
		return fmt.Sprintf("%v is up for maintenance", name), nil
	} else {
		return fmt.Sprintf("%v is all good", name), nil
	}
}
