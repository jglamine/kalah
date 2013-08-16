package kalah

import (
	"testing"
	)

func TestMakeGameNode(t* testing.T) {
	cells := []byte{2, 1, 0, 6, 8, 11, 4,
					 3, 1, 0, 0, 0, 9, 3}
	b := MakeBoardFromCells(PlayerTwo, cells)
	parentState := boardToState(b)
	if parentState.LegalMove(byte(10)) {
		t.Errorf("10 should not be a legal move")
	}
	makeGameNode(byte(7), parentState)
	makeGameNode(byte(8), parentState)
	if parentState.LegalMove(byte(9)) {
		t.Errorf("9 should not be a legal move")
	}
	if parentState.LegalMove(byte(10)) {
		t.Errorf("10 should not be a legal move")
	}
}

func TestMakeChildren(t *testing.T) {
	cells := []byte{2, 1, 0, 6, 8, 11, 4,
					 3, 1, 0, 0, 0, 9, 3}
	b := MakeBoardFromCells(PlayerTwo, cells)
	state := boardToState(b)
	children := makeChildren(state)
	for _, child := range children {
		if child.state.WhoseTurn() != PlayerOne {
			t.Errorf("Child node (move %v) has wrong state.WhoseTurn()",
				child.move)
		}
	}
	if len(children) != 3 {
		t.Errorf("Wrong number of children: expected: 3, got: %v\n",
			len(children))
		for _, child := range children {
			t.Errorf("child move: %v", child.move)
		}
	}
}

// TODO: test minimax algorithm for correctness