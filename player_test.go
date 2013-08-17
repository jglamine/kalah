package kalah

import (
	"testing"
	)

func TestBozo(t *testing.T) {
	cells := []byte{4, 4, 0, 5, 5, 5, 1,
					4, 4, 4, 4, 4, 4, 0}
	b := MakeBoardFromCells(PlayerOne, cells)
	bozo := MakeBozoPlayer("Bozo")
	move, _ := bozo.ChooseMove(b)
	if move != byte(5) {
		t.Error("Move test 1: expected 5, got", move)
	}
}