/*
	The strategy pattern (also known as the policy pattern) is a behavioral software design pattern that enables selecting an algorithm at runtime. Instead of implementing a single algorithm directly, code receives run-time instructions as to which in a family of algorithms to use.
*/

package designPattern

import "fmt"

type Vehicle struct {
	brakeSystem BrakeBehavior
	brandName   string
}

func NewVehicle(brakeSystem BrakeBehavior, brandName string) *Vehicle {
	return &Vehicle{brakeSystem: brakeSystem, brandName: brandName}
}

func (v Vehicle) ApplyBrake() {
	v.brakeSystem.Brake()
}

func (v Vehicle) ShowBrand() {
	fmt.Println("I'm ", v.brandName)
}

type BrakeBehavior interface {
	Brake()
}

type Brake struct{}

func (b Brake) Brake() {
	fmt.Println("braking until fully stopped")
}

type ABSBrake struct{}

func (A ABSBrake) Brake() {
	fmt.Println("brake and release and brake")
}

type ElectricBrake struct{}

func (e ElectricBrake) Brake() {
	fmt.Println("powered by electrics instead of hydraulic")
}

func StrategyDesignDemo() {
	bmw := NewVehicle(ABSBrake{}, "bmw")
	bmw.ApplyBrake()

	honda := NewVehicle(Brake{}, "honda")
	honda.ApplyBrake()

	alfaRomeo := NewVehicle(ElectricBrake{}, "alfaRomeo")
	alfaRomeo.ApplyBrake()
}
