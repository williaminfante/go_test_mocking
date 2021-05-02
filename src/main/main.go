package main

import (
	"fmt"

	"github.com/williaminfante/go_test_mocking/src/maintenance"
	"github.com/williaminfante/go_test_mocking/src/vehicle"
)

func main() {
	realSpecialist := new(maintenance.Specialist)
	maintenanceRepository := vehicle.MaintenanceWrapper{Maintenance: realSpecialist}
	fmt.Println(maintenanceRepository.ConductChecks("lightning", 45))
	fmt.Println(maintenanceRepository.ConductChecks("lightning", 30))
	fmt.Println(maintenanceRepository.ConductChecks("lightning", -24))
}
