package kalah

import (
	"testing"
	"time"
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

// TestConsistent makes sure that minimax consistently gives the same result
// when playing on different halves of the board.
func TestConsistent(t* testing.T) {
	p1 := MakeMinimaxPlayer("Minimax (depth 3)", time.Duration(120)*time.Second, 3)
	p2 := MakeMinimaxPlayer("Minimax (depth 1)", time.Duration(60)*time.Millisecond, 1)
	p1Score1, p2Score1 := RunMatch(p1, p2, PlayerOne, false)
	p2Score2, p1Score2 := RunMatch(p2, p1, PlayerTwo, false)
	if p1Score1 != p1Score2 || p2Score1 != p2Score2 {
		t.Errorf("Scores do not match:\ngame 1: %v to %v\ngame2: %v to %v",
			p1Score1, p2Score1, p1Score2, p2Score2)
	}
	if p1Score1 <= p2Score1 || p1Score2 <= p2Score2 {
		t.Errorf("%v lost to %v\ngame 1: %v to %v\ngame 2: %v to %v",
			p1, p2, p1Score1, p2Score1, p1Score2, p2Score2)
	}
}
// TODO: test minimax algorithm for correctness
