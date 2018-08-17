package main

import (
	"fmt"
)

type Pizza struct {
	Name string
}

type IPizza interface {
	Describe(func(pz *Pizza)) func(pz *Pizza)
}

func (pizza Pizza) Describe() func(pz *Pizza) {
	return func(pz *Pizza) {
		pz = &pizza
		fmt.Printf("Pizza name: %s\n", pizza.Name)
	}
}

type BehaviorDecorator struct{}

func (d BehaviorDecorator) Describe(f func(pz *Pizza)) func(pz *Pizza) {
	return func(pz *Pizza) {
		f(pz)
		pz.Name = "Behavior Name"
		fmt.Printf("I changed pizza's name to: %s\n", pz.Name)
	}
}

type StateDecorator struct {
	hot string
}

func (stateDecorator StateDecorator) Describe(f func(pz *Pizza)) func(pz *Pizza) {
	return func(pz *Pizza) {
		f(pz)
		fmt.Printf("Added hot string: %s\n", stateDecorator.hot)
	}
}

func main() {
	fmt.Println("Decorator Pattern Demo")
	pizza := Pizza{Name: "Original Name"}
	pizza.Describe()(&pizza)
	newBehavior := BehaviorDecorator{}
	newState := StateDecorator{hot: "so hot"}
	newBehavior.Describe(newState.Describe(pizza.Describe()))(&pizza)
}
