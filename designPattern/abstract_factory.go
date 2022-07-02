/*
	Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
	Use when you have many objects that can be added or changed dynamically during runtime.
	You can model anything you can imagine and have those objects interact through common interfaces
*/

package designPattern

import "fmt"

type transportationToolType = int

const (
	TransportationAirplane transportationToolType = iota + 1
	TransportationTrain
)

type TransportationBureau struct{}

func NewTransportationBureau() *TransportationBureau {
	return &TransportationBureau{}
}

func (tb *TransportationBureau) BuildFactory(transportationTool transportationToolType) TransportationFactory {
	switch transportationTool {
	case TransportationAirplane:
		return NewAirplaneFactory()
	case TransportationTrain:
		return NewTrainFactory()
	}
	return NewAirplaneFactory()
}

type TransportationFactory interface {
	ProduceTransportationTools(name string) TransportationTool
}

type AirplaneFactory struct{}

func NewAirplaneFactory() *AirplaneFactory {
	return &AirplaneFactory{}
}

func (a *AirplaneFactory) ProduceTransportationTools(name string) TransportationTool {
	return NewAirplane(name)
}

type TrainFactory struct{}

func NewTrainFactory() *TrainFactory {
	return &TrainFactory{}
}

func (t *TrainFactory) ProduceTransportationTools(name string) TransportationTool {
	return NewTrain(name)
}

type TransportationTool interface {
	Speed()
	Cost()
	Name()
	Size()
}

type Airplane struct {
	name string
}

func NewAirplane(name string) *Airplane {
	return &Airplane{name: name}
}

func (a *Airplane) Speed() {
	fmt.Println("fast as like sounds")
}

func (a *Airplane) Cost() {
	fmt.Println("only billionaires could afford")
}

func (a *Airplane) Name() {
	fmt.Println("I'm ", a.name)
}

func (a *Airplane) Size() {
	fmt.Println("usually takes hundred peoples")
}

type Train struct {
	name string
}

func NewTrain(name string) *Train {
	return &Train{name: name}
}

func (t Train) Speed() {
	fmt.Println("not very fast but much more than most cars")
}

func (t Train) Cost() {
	fmt.Println("just companies or government could afford")
}

func (t Train) Name() {
	fmt.Println("I'm ", t.name)
}

func (t Train) Size() {
	fmt.Println("could take thousand peoples")
}

func AbstractFactoryDesignDemo() {
	transportationBureau := NewTransportationBureau()
	airPlanesFactory, trainsFactory := transportationBureau.BuildFactory(TransportationAirplane), transportationBureau.BuildFactory(TransportationTrain)

	boeing := airPlanesFactory.ProduceTransportationTools("Boeing from US")
	boeing.Name()
	boeing.Size()
	boeing.Cost()
	boeing.Speed()

	fmt.Println()

	bayStream := airPlanesFactory.ProduceTransportationTools("Bay Stream from CA")
	bayStream.Name()
	bayStream.Size()
	bayStream.Cost()
	bayStream.Speed()

	fmt.Println()

	concord := trainsFactory.ProduceTransportationTools("协和号 from CN")
	concord.Name()
	concord.Size()
	concord.Cost()
	concord.Speed()
}
