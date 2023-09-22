package sound

import "fmt"

type Dog struct {
	Breed string
}

func (dog *Dog) Speak() string {
	return fmt.Sprintf("Woof! My breed is %s!", dog.Breed)
}
