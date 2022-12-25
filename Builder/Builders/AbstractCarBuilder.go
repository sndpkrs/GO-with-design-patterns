package builders

import "fmt"

type Car struct {
	frame     string
	engine    string
	steering  string
	noOfSeats int
	isRoofTop bool
}

func (car *Car) Drive() {
	fmt.Println("Driving ", car.frame, car.engine, car.steering, car.noOfSeats, car.isRoofTop)
}

type AbstractCarBuilder interface {
	BuildFrame() AbstractCarBuilder
	BuildEngine() AbstractCarBuilder
	BuildSteering() AbstractCarBuilder
	BuildNoOfSeats() AbstractCarBuilder
	BuildRoofTop() AbstractCarBuilder

	Build() Car
}
