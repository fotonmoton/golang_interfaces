package main

import (
	"errors"
	"interfaces/cache"
	cruelerror "interfaces/cruelError"
	"interfaces/decorator"
	racegame "interfaces/raceGame"
	"interfaces/sound"
	"log"
	"os"
	"strings"
)

/*
Interfaces
0. Receivers can be on different types, not only structs
1. Set of method signatures https://go.dev/tour/methods/9
2. It's a type of polymorphism
3. Also named as protocols in another languages
4. We can define variable of specific interface, can accept it as argument, etc.
5. Pointer receivers and value receivers are different implementations
6. Implemented implicitly, no need to specify interface for implementation https://go.dev/tour/methods/10
7. Interfaces could have nil values but concrete type https://go.dev/tour/methods/12
8. Interfaces as tuple (value, type) https://go.dev/tour/methods/11
9. "nil interface values" https://go.dev/tour/methods/13
10. Type must implement all methods to satisfy interface
11. Type can implement multiple interfaces
12. Empty interface interface{} https://go.dev/tour/methods/14
13. Keep interfaces small
14. interface embedding, in structs too, partial satisfaction
15. type assertion https://go.dev/tour/methods/15
16. anonymous interfaces
17. can't be satisfied if partially exported


https://go.dev/tour/methods/9 to https://go.dev/tour/methods/23
*/

func ErrorPrinter(err error) {
	log.Println(err.Error())
}

func Animals() {
	var dog sound.Speakable = &sound.Dog{Breed: "Jack Russel terrier"}

	var mic sound.Playable = &sound.Microphone{}

	mic.Play(dog)

	cat := &sound.Cat{Lives: 9}

	mic.Play(cat)
}

func Errors() {
	ErrorPrinter(errors.New("standard error"))
	ErrorPrinter(new(cruelerror.CruelError))
}

func Cache() {
	lru := cache.NewCache(&cache.Lru{})

	lru.Add("hello", "world")

	lru.Add("another", "world")

	fifio := cache.NewCache(&cache.Fifo{})

	fifio.Add("hello", "world")

	fifio.Add("another", "world")
}

func Decorator() {
	itself, err := os.Open("main.go")

	if err != nil {
		log.Fatal(err)
	}

	defer itself.Close()

	decorator.Decorate(itself, "*")

	decorator.Decorate(strings.NewReader("Simple string"), "-")

}

func RaceGame() {
	players := map[racegame.Drivable]int{
		&racegame.SportCar{MaxSpeed: 120, Brand: "BMW"}:    0,
		&racegame.Horse{Name: "Lucky", MaxRideWeight: 100}: 0,
	}

	racegame.RaceGame(players)
}

func main() {
	Animals()
	Errors()
	Cache()
	Decorator()
	RaceGame()
}
