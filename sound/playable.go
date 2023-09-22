package sound

// This interface allows to play some strings
// It returns how many letters from string was played
type Playable interface {
	Play(Speakable) int
}
