package kalah

// Player is an interface for a kalah player.
// To implement the Player interface, subtype player and implement ChooseMove()
type Player interface {
	ChooseMove(b Board) byte
	String() string
}

// player is an abstract base player type
// It does not fully implement the Player interface.
// To implement the Player interface, subtype player and implement ChooseMove()
type player struct {
	name string
}

// bozoPlayer is a simple AI which implements Player
type bozoPlayer struct {
	player
}

func MakeBozoPlayer(name string) Player {
	player := bozoPlayer{player{name}}
	return &player
}

func (this *player) String() string {
	return this.name
}

// ChooseMove is bozo's basic kalah AI
// It tries to chooses a move which lets it go again, defaulting the the first
// available move.
func (this *bozoPlayer) ChooseMove(b Board) byte {
	if b.WhoseTurn() == PlayerOne {
		// try to find a go-again move
		for i := kalahOne-1; i != byte(255); i-- {
			if b.Cell(i) == kalahOne-i { return i }
		}
		// otherwise, use first available move
		for i := kalahOne-1; i != byte(255); i-- {
			if b.Cell(i) > byte(0) { return i }
		}
	} else {
		// try to find a go-again move
		for i := kalahTwo-1; i > kalahOne; i-- {
			if b.Cell(i) == kalahTwo-i { return i }
		}
		for i := kalahTwo-1; i > kalahOne; i-- {
			if b.Cell(i) > byte(0) { return i }
		}
	}
	// if no legal moves are found, return an illegal one
	// this can not happen unless the game is over
	return kalahOne
}
