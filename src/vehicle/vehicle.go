package vehicle

import (
	"fmt"
)

type MaintenanceAPI interface {
	NeedsMaintenance(days int) (decision bool, err error)
}
type MaintenanceWrapper struct {
	Maintenance MaintenanceAPI
}

func (w *MaintenanceWrapper) ConductChecks(name string, days_used int) (string, error) {
	decision, errMaintenance := w.Maintenance.NeedsMaintenance(days_used)
	if errMaintenance != nil {
		return "", errMaintenance
	}
	if decision {
		return fmt.Sprintf("%v is up for maintenance", name), nil
	} else {
		return fmt.Sprintf("%v is all good", name), nil
	}
}
