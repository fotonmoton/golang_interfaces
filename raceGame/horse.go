package racegame

import (
	"fmt"
	"math/rand"
)

type Horse struct {
	Name          string
	MaxRideWeight float32
}

func (horse *Horse) Drive() int {
	meters := rand.Intn(10) + 1
	horse.MaxRideWeight = horse.MaxRideWeight - 2
	fmt.Printf("Horse rides for %d meters for this step\n", meters)
	return meters
}
