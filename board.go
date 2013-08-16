package kalah

import (
	"strconv"
	)

type playerId byte;

const(
	// number of stones per cell
	numStones = byte(4)
	// number of cells
	numCells = byte(14)
	kalahOne = (numCells -1) / 2 // 6
	kalahTwo = numCells -1 // 13
	PlayerOne = playerId(0)
	PlayerTwo = playerId(1)
	PlayerTie = playerId(2)
)

/* Board is a mankalah board.
*/
type Board interface {
	LegalMove(cell byte) bool
	Cell(cell byte) byte
	WhoseTurn() playerId
	GameOver() bool
	Winner() playerId
	Score() (byte, byte)
	Move(cell byte) (numCaptured byte, ok bool)
	String() string
}

type board struct {
	playerToMove playerId
	cells []byte
}

func (b board) LegalMove(cell byte) bool {
	if b.playerToMove == PlayerOne {
		if cell >= byte(0) && cell < kalahOne && b.cells[cell] != byte(0) {
			return true
		}
		return false
	}
	if cell > kalahOne && cell < kalahTwo && b.cells[cell] != byte(0) {
		return true
	}
	return false
}

func (b board) Cell(cell byte) byte {
	return b.cells[cell]
}

func (b board) WhoseTurn() playerId {
	return b.playerToMove
}

func (b board) GameOver() bool{
	return (allEmpty(b.cells[0:kalahOne]) ||
		   allEmpty(b.cells[kalahOne+1:kalahTwo]))
}

func (b board) Winner() playerId {
	score1, score2 := b.Score()
	if score1 > score2 {
		return PlayerOne
	}
	if score2 > score1 {
		return PlayerTwo
	}
	return PlayerTie
}

func (b board) Score() (p1Score, p2Score byte) {
	p1Score = byte(0)
	p2Score = byte(0)
	for _, cell := range b.cells[0:kalahOne+1] {
		p1Score += cell
	}
	for _, cell := range b.cells[kalahOne+1:kalahTwo+1] {
		p2Score += cell
	}
	return
}

// Move executes a move in a cell for the current player
// The number of stones captured and an error flag are returned.
func (b *board) Move(cell byte) (numCaptured byte, ok bool) {
	numCaptured = byte(0)
	if !b.LegalMove(cell) {
		return numCaptured, false
	}
	// pick up the stones
	stones := b.cells[cell]
	b.cells[cell] = byte(0)
	// distribute stones around the board
	pos := cell+1
	for ; stones > byte(0); pos++ {
		// don't add stone to opponents kalah
		if b.playerToMove == PlayerOne && pos == kalahTwo {
			pos++
		} else if b.playerToMove == PlayerTwo && pos == kalahOne {
			pos++
		}
		if pos == numCells {
			pos = byte(0)
		}
		b.cells[pos]++
		stones--
	}
	pos--
	// if possible, capture stones on opponents side
	// capture is allowed when ending on an empty cell on your side
	if b.playerToMove == PlayerOne && pos < kalahOne &&
		b.cells[pos] == 1 && b.cells[kalahTwo-1-pos] > 0 {
	   	numCaptured = b.cells[kalahTwo-1-pos] + 1
	   	b.cells[kalahOne] += b.cells[kalahTwo-1-pos]
	   	b.cells[kalahTwo-1-pos] = 0
	   	b.cells[kalahOne]++
	   	b.cells[pos] = 0
	} else if (b.playerToMove == PlayerTwo && pos > kalahOne && pos < kalahTwo &&
		b.cells[pos] == 1 && b.cells[kalahTwo-1-pos] > 0) {
		numCaptured = b.cells[kalahTwo-1-pos] + 1
      	b.cells[kalahTwo] += b.cells[kalahTwo-1-pos]
      	b.cells[kalahTwo-1-pos] = 0
      	b.cells[kalahTwo]++
      	b.cells[pos] = 0
	}
	// if player doesn't get to go again, change who's turn it is
	if b.playerToMove == PlayerTwo {
		if pos != kalahTwo {
			b.playerToMove = PlayerOne
		}
	} else if pos != kalahOne {
		b.playerToMove = PlayerTwo
	}
	return numCaptured, true
}

func (b board) String() string {
	s := "   "
	for i := kalahTwo-1; i > kalahOne; i-- {
		s += strconv.Itoa(int(b.cells[i])) + "  "
	}
	s += "\n" + strconv.Itoa(int(b.cells[kalahTwo])) + "  "
	for i := byte(0); i < kalahOne; i++ {
		s += "   "
	}
	s += strconv.Itoa(int(b.cells[kalahOne])) + "\n   "
	for i := byte(0); i < kalahOne; i++ {
		s += strconv.Itoa(int(b.cells[i])) + "  "
	}
	return s
}

func MakeBoard(playerToMove playerId) Board {
	b := new(board)
	b.playerToMove = playerToMove
	b.cells = make([]byte, numCells)
	for i := range b.cells {
		b.cells[i] = numStones
	}
	b.cells[kalahOne] = 0
	b.cells[kalahTwo] = 0
	return b
}

func MakeBoardFromCells(playerToMove playerId, cells []byte) Board {
	b := new(board)
	b.playerToMove = playerToMove
	b.cells = cells
	return b
}

// allEmpty returns true if all the cells are empty.
func allEmpty(cells []byte) bool {
	for _, cell := range cells {
		if cell != 0 {
			return false
		}
	}
	return true
}
