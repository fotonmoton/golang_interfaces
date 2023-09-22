package racegame

// This interface defines drive behavior
// It returns how many meters entity drove
type Drivable interface {
	Drive() int
}

// Returns true if really stopped
type Stopable interface {
	Stop() bool
}

type Racable interface {
	Drivable
	Stopable
}
