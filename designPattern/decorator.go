package designPattern

import "fmt"

/*
	Usually used for refactor
*/

type VehicleDecorator struct {
	Vehicle
}

func NewVehicleDecorator(vehicle Vehicle) *VehicleDecorator {
	return &VehicleDecorator{Vehicle: vehicle}
}

// ShowBrand Override Vehicle.ShowBrand
func (vd VehicleDecorator) ShowBrand() {
	vd.Vehicle.ShowBrand()
	fmt.Println("with Bose Speaker")
}

func DecoratorDesignDemo() {
	vehicleDecorator := NewVehicleDecorator(*NewVehicle(ABSBrake{}, "bmw"))
	vehicleDecorator.ApplyBrake()
	vehicleDecorator.ShowBrand()
}
