package main

import "fmt"

type Observer interface {
	Update(*ConcreteSubject)
}

type NameObserver struct {
	Name string
}

func (nameObs *NameObserver) Update(subject *ConcreteSubject) {
	fmt.Printf("Updated Name from: %s to %s\n", nameObs.Name, subject.Name)
	nameObs.Name = subject.Name
}

type PriceObserver struct {
	Price int
}

func (priceObs *PriceObserver) Update(subject *ConcreteSubject) {
	fmt.Printf("Updated Price from: %d to %d\n", priceObs.Price, subject.Price)
	priceObs.Price = subject.Price
}

type Subject interface {
	Notify()
	Attach(Observer)
	Detach(Observer)
}

// An event
type ConcreteSubject struct {
	Name         string
	Price        int
	ObserversMap map[Observer]struct{}
}

func (subject *ConcreteSubject) makeChanges(name string, price int) {
	subject.Name = name
	subject.Price = price
}

func (subject *ConcreteSubject) Attach(obs Observer) {
	subject.ObserversMap[obs] = struct{}{}
}

func (subject *ConcreteSubject) Detach(obs Observer) {
	delete(subject.ObserversMap, obs)
}

func (subject ConcreteSubject) Notify() {
	for obs := range subject.ObserversMap {
		obs.Update(&subject)
	}
}

func main() {
	fmt.Println("Observer Pattern Demo")
	newEvent := ConcreteSubject{Name: "Orgin",
		Price:        1,
		ObserversMap: map[Observer]struct{}{}}
	newNameObserver := NameObserver{Name: "Blank"}
	newPriceObserver := PriceObserver{Price: 0}
	newEvent.Attach(&newNameObserver)
	newEvent.Attach(&newPriceObserver)
	newEvent.Notify()
	newEvent.makeChanges("new name", 100)
	newEvent.Notify()
	newEvent.makeChanges("another name", -1)
	newEvent.Notify()
}
