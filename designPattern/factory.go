package designPattern

/*
	Also called Factory Method
	Define an interface for creating an object but let the classes that implement the interface decide which class to instantiate.
	The factory method lets a class defer instantiation to subclasses.
*/

type VehicleFactory struct{}

func NewVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

func (vf *VehicleFactory) ProduceVehicle(brand string) *Vehicle {
	switch brand {
	case "bmw":
		return NewVehicle(ABSBrake{}, "bmw")
	case "honda":
		return NewVehicle(Brake{}, "honda")
	case "alfaRomeo":
		return NewVehicle(ElectricBrake{}, "alfaRomeo")
	}
	return NewVehicle(Brake{}, "no brand")
}

func FactoryDesignDemo() {
	vehicleFactory := NewVehicleFactory()
	bmw := vehicleFactory.ProduceVehicle("bmw")
	bmw.ShowBrand()
	honda := vehicleFactory.ProduceVehicle("honda")
	honda.ShowBrand()
	alfaRomeo := vehicleFactory.ProduceVehicle("alfaRomeo")
	alfaRomeo.ShowBrand()
}
