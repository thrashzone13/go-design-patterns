package factory

import "fmt"

func CreateVehicle(vehicleType string, name string, color string) (Vehicler, error) {
	switch vehicleType {
	case "car":
		return createCar(name, color, 100), nil
	case "truck":
		return createTruck(name, color, 18), nil
	default:
		return nil, fmt.Errorf("unknown vehicle type")
	}
}
