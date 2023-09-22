package racegame

import (
	"fmt"
	"math/rand"
)

type SportCar struct {
	MaxSpeed int
	Brand    string
}

func (car *SportCar) Drive() int {
	meters := rand.Intn(10) + 5
	fmt.Printf("Sport car drives for %d meters for this step\n", meters)
	return meters
}

func (car *SportCar) Stop() bool {
	return true
}

func (car *SportCar) Tune() *SportCar {
	car.MaxSpeed += 10
	return car
}
