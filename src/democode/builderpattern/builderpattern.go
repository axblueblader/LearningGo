package main

/*	BUILDER PATTERN
Purpose:
	+ Encapsulates the construction logic for complex object
	+ With the same process, create different presentations
Components:
	+ Product : data model for main product
	+ Abstract Builder : an interface for different builders
	+ Concrete Builder : where the real construction happens and provides
	interface for retrieving final product
	+ Director : is essentially an assembly line where we give it
	the builder for type of product we want and construct product through
	it's interface
Different implementations:
	Type 1:
		assembly = newBuilder(UniqueBuilder)
		product = assembly.buildProduct()
	Type 2:
		assembly = newBuilder().setBuilderProperties()
		product = assembly.Part1(part1).Part2(part2).Build()
*/

import (
	"fmt"
)

//Name : name of product
type Name string

//Price : price of product
type Price int

//Description : description of product
type Description string

//Product : our final wanted product
type Product struct {
	name        Name
	price       Price
	description Description
}

//Show : display product information
func (product Product) Show() {
	fmt.Printf("Name: %v\nPrice: %v\nDescription: %v\n",
		product.name,
		product.price,
		product.description)
}

//AbstractBuilder : the abstract builder
type AbstractBuilder interface {
	Build() Product
}

//WindowsBuilder : concrete builder for Windows
type WindowsBuilder struct{}

//Build (WindowsBuilder) : create a Windows PC
func (compBuilder WindowsBuilder) Build() Product {
	computer := Product{
		name:        "Windows",
		price:       100,
		description: "This model was built to run on Windows"}
	return computer
}

//MacBuilder : concrete builder for Mac
type MacBuilder struct{}

//Build (MacBuilder) : create a Mac
func (macBuilder MacBuilder) Build() Product {
	computer := Product{
		name:        "OS X",
		price:       200,
		description: "This model was built to run on OS X"}
	return computer
}

//Director : make product through builder interface
//can be regarded as an assembly line
type Director struct {
	builder AbstractBuilder
}

//newBuilder : users provide the Director which builder they want to use
func (director *Director) newBuilder(concreteBuilder AbstractBuilder) {
	director.builder = concreteBuilder
}

//makeComputer : an interface for users to build product they want
//without needing to care about how it is built
func (director *Director) makeComputer() Product {
	return director.builder.Build()
}

func main() {
	fmt.Println("Builder Pattern Demo")
	var assembly Director
	assembly.newBuilder(WindowsBuilder{})
	product := assembly.makeComputer()
	product.Show()
	assembly.newBuilder(MacBuilder{})
	product = assembly.makeComputer()
	product.Show()
}
