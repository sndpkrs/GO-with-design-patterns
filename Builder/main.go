package main

import (
	builders "builder/Builders"
	"fmt"
)

func main() {
	builder := getBuilder("bmw")
	gd := NewGarageDirector(builder)
	car := gd.BuildCar()
	fmt.Println(car)
}

func getBuilder(carType string) builders.AbstractCarBuilder {
	if carType == "bmw" {
		return builders.NewBMWBuilder()
	} else {
		return builders.NewMiniBuilder()
	}
}
