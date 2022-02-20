package game

// Game represents an abstract game that tracks a running score.
type Game interface {
	// Begin is used to identify when the game has begun.
	Begin() error
}
