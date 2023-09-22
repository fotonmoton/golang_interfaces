package sound

import "fmt"

type Cat struct {
	Lives int
}

func (cat *Cat) Speak() string {
	return fmt.Sprintf("Meow! I left with %d lives!", cat.Lives)
}
