package sound

import (
	"fmt"
)

type Microphone struct{}

func (m *Microphone) Play(s Speakable) int {
	words := s.Speak()

	fmt.Println(words)

	return len(words)
}
