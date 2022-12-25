package main

import builders "builder/Builders"

type GarageDirector struct {
	builder builders.AbstractCarBuilder
}

func NewGarageDirector(acb builders.AbstractCarBuilder) *GarageDirector {
	return &GarageDirector{
		builder: acb,
	}
}

func (gd *GarageDirector) BuildCar() builders.Car {
	gd.builder.BuildFrame().BuildEngine().BuildSteering().BuildNoOfSeats().BuildRoofTop()
	return gd.builder.Build()
}
