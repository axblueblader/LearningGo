package main

import (
	"fmt"
)

type Pizza interface {
	Describe(Pz Pizza)
	SetName(name string)
	GetName() string
}

type VietnamesePizza struct {
	Name    string
	EggType string
}

func (pz *VietnamesePizza) Describe(Pz Pizza) {
	fmt.Printf("Different pizza type: %s with a different member: %s\n",
		pz.Name,
		pz.EggType)
}

func (pz *VietnamesePizza) SetName(name string) {
	pz.Name = name
}

func (pz *VietnamesePizza) GetName() string {
	return pz.Name
}

type ItalianPizza struct {
	Name string
}

func (pz *ItalianPizza) Describe(Pz Pizza) {
	fmt.Printf("My pizza type: %s\n", pz.Name)
}

func (pz *ItalianPizza) SetName(name string) {
	pz.Name = name
}

func (pz *ItalianPizza) GetName() string {
	return pz.Name
}

type Decorator interface {
	Decorate(func(pz Pizza)) func(pz Pizza)
}

type BehaviorDecorator struct{}

func (d BehaviorDecorator) Decorate(f func(pz Pizza)) func(pz Pizza) {
	return func(pz Pizza) {
		f(pz)
		pz.SetName("Behavior Name")
		fmt.Printf("I changed pizza's name to: %s\n", pz.GetName())
	}
}

type StateDecorator struct {
	hot string
}

func (stateDecorator StateDecorator) Decorate(f func(pz Pizza)) func(pz Pizza) {
	return func(pz Pizza) {
		f(pz)
		fmt.Printf("Added hot string: %s to %s\n", stateDecorator.hot, pz.GetName())
	}
}

func main() {
	fmt.Println("Decorator Pattern Demo")
	pizza := ItalianPizza{Name: "Original Itatalian Name"}
	newBehavior := BehaviorDecorator{}
	newState := StateDecorator{hot: "so hot"}
	decorator := newBehavior.Decorate(newState.Decorate(pizza.Describe))
	decorator(&pizza)
	vietPizza := VietnamesePizza{Name: "Banh trang nuong", EggType: "Trung cut"}
	decorator = newState.Decorate(vietPizza.Describe)
	decorator(&vietPizza)
}
