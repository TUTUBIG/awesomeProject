package designPattern

import "fmt"

/*
	The strategy pattern (also known as the policy pattern) is a behavioral software design pattern that enables selecting an algorithm at runtime. Instead of implementing a single algorithm directly, code receives run-time instructions as to which in a family of algorithms to use.
*/

type Vehicle struct {
	brakeSystem BrakeBehavior
}

func (v Vehicle) ApplyBrake() {
	v.brakeSystem.Brake()
}

func CreateVehicle(brakeSystem BrakeBehavior) Vehicle {
	return Vehicle{brakeSystem: brakeSystem}
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
	bmw := CreateVehicle(ABSBrake{})
	bmw.ApplyBrake()

	honda := CreateVehicle(Brake{})
	honda.ApplyBrake()

	alfaRomeo := CreateVehicle(ElectricBrake{})
	alfaRomeo.ApplyBrake()
}
