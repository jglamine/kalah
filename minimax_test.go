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
// when playing on different halves of the board. It also checks that it is
// able to beat randomly chosen moves.
func TestConsistent(t* testing.T) {
	p1 := MakeMinimaxPlayer("Minimax (depth 3)", time.Duration(120)*time.Second, 3)
	p2 := MakeMinimaxPlayer("Minimax (depth 1)", time.Duration(120)*time.Second, 1)
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

// BenchmarkMinimax measures how quickly minimax runs
func BenchmarkMinimaxMatch(b *testing.B) {
    for i := 0; i < b.N; i++ {
        p1 := MakeMinimaxPlayer("Minimax (depth 7)", time.Duration(120)*time.Second, 7)
		p2 := MakeMinimaxPlayer("Random", time.Duration(0), 0)
		RunMatch(p1, p2, PlayerOne, false)
    }
}

// BenchmarkSearch measures the speed of a single minimax search
func BenchmarkSearch(b *testing.B) {
	depthLimit := 10
	tree := makeGameTree(MakeBoard(PlayerTwo))
	for i := 0; i < b.N; i++ {
		tree.search(depthLimit, time.Now().Add(time.Duration(120)*time.Second))
	}
}