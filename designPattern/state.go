package designPattern

import "fmt"

type Light struct {
	State LightState
}

func (l *Light) ClickWitchButton() {
	l.State.SwitchState(l)
}

func NewLight() *Light {
	return &Light{State: &LightOff{}}
}

type LightState interface {
	SwitchState(light *Light)
	describe() string
}

type LightOn struct{}

func (l LightOn) describe() string {
	return "on"
}

func (l LightOn) SwitchState(light *Light) {
	if light.State.describe() != "on" {
		fmt.Println("invalid state")
		return

	}

	fmt.Println("switch off the light")
	light.State = &LightOff{}
}

type LightOff struct{}

func (l LightOff) describe() string {
	return "off"
}

func (l LightOff) SwitchState(light *Light) {
	if light.State.describe() != "off" {
		fmt.Println("invalid state")
		return
	}

	fmt.Println("switch on the light")
	light.State = &LightOn{}
}

func StateDesignDemo() {
	light := NewLight()

	light.ClickWitchButton()
	light.ClickWitchButton()
	light.ClickWitchButton()
	light.ClickWitchButton()
}
