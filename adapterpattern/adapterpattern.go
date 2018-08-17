package main

import "fmt"

type Car interface {
	Run()
}

type NormalCar struct {
	Name string
}

func (car NormalCar) Run() {
	fmt.Println("Normal car is running")
}

type RaceCar struct {
	Name string
}

func (car RaceCar) Run() {
	fmt.Println("Race car is speeding")
}

type FlyingCar struct {
	Name string
}

func (car FlyingCar) Drive() {
	fmt.Println("Flying car is driving")
}

func (car FlyingCar) Fly() {
	fmt.Println("Flying car is flying")
}

type AdvancedCar struct {
	flyCar FlyingCar
	Name   string
}

func (advCar AdvancedCar) Run() {
	advCar.flyCar.Drive()
}

type Client struct {
	car Car
}

func (client Client) makeCarRun() {
	client.car.Run()
}

func main() {
	fmt.Println("Adapter Pattern Demo")
	client := Client{car: NormalCar{Name: "Normal Car"}}
	client.makeCarRun()
	client.car = RaceCar{Name: "Race Car"}
	client.makeCarRun()
	client.car = AdvancedCar{Name: "Advanced Car"}
	client.makeCarRun()
}
