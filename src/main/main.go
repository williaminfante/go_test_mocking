package main

import (
	"fmt"

	"github.com/williaminfante/go_test_mocking/src/vehicle"
)

func main() {
	fmt.Println(vehicle.ConductChecks("lightning", 45))
	fmt.Println(vehicle.ConductChecks("lightning", 30))
	fmt.Println(vehicle.ConductChecks("lightning", -24))
}
