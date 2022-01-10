package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Manufacture
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() Vehicle
}

type Builder struct{
	process BuildProcess
}

func (b *Builder) Construct() {
	b.process.SetWheels().SetSeats().SetStructure()
}

func (b *Builder) SetBuilder(bp BuildProcess) {
	b.process = bp
}

type Vehicle struct {
	Wheels int
	Seats int
	Structure string
}

// Car ------------------------------------------------
type CarBuilder struct {
	car Vehicle
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.car.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.car.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.car.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() Vehicle {
	return c.car
}

// Bike -----------------------------------------------
type BikeBuilder struct {
	bike Vehicle
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.bike.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.bike.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.bike.Structure = "bike"
	return b
}

func (b *BikeBuilder) GetVehicle() Vehicle {
	return b.bike
}

// Main ----------------------------------------------
// Строитель — это порождающий паттерн проектирования, который позволяет создавать объекты пошагово.
func main() {
	builder := Builder{}

	carBuilder := CarBuilder{}

	builder.SetBuilder(&carBuilder)
	builder.Construct()

	car := carBuilder.GetVehicle()
	fmt.Println(car)

	bikeBuilder := BikeBuilder{}
	builder.SetBuilder(&bikeBuilder)
	builder.Construct()

	bike := bikeBuilder.GetVehicle()
	fmt.Println(bike)
}