package designPattern

import "fmt"

type Chopsticks interface {
	Pickup()
}

type PublicChopsticks struct{}

func NewPublicChopsticks() *PublicChopsticks {
	return &PublicChopsticks{}
}

func (p PublicChopsticks) Pickup() {
	fmt.Println("pick up foods from public dishes to our own, then use our private chopsticks to eat food")
}

type PrivateChopsticks struct {
	owner string
}

func NewPrivateChopsticks(owner string) *PrivateChopsticks {
	return &PrivateChopsticks{owner: owner}
}

func (p PrivateChopsticks) Pickup() {
	fmt.Printf("pick up foods from %s's own dish to eat\n", p.owner)
}

type Person struct {
	privateChopsticks Chopsticks
}

func NewPerson(privateChopsticks Chopsticks) *Person {
	return &Person{privateChopsticks: privateChopsticks}
}

func (p Person) EatFood(publicChopsticks Chopsticks) {
	publicChopsticks.Pickup()
	p.privateChopsticks.Pickup()
}

func SingletonDesignDemo() {
	publicChopsticks := NewPublicChopsticks()

	michael := NewPerson(NewPrivateChopsticks("michael"))
	michelle := NewPerson(NewPrivateChopsticks("michelle"))

	michelle.EatFood(publicChopsticks)
	michael.EatFood(publicChopsticks)
}
