package kalah

import (
	"testing"
	)

func TestMakeBoard(t *testing.T) {
	b := MakeBoard(PlayerOne)
	if b.WhoseTurn() != PlayerOne {
		t.Fatal("playerToMove != PlayerOne")
	}
	b2 := MakeBoard(PlayerTwo)
	if b2.WhoseTurn() != PlayerTwo {
		t.Fatal("playerToMove != PlayerTwo")
	}
	if numStones != 4 || numCells != 14 {
		t.Fatal("Constants numCells and NumStones are wrong")
	}
	if b.Cell(6) != 0 {
		t.Fatal("Board cell 6 is not empty")
	}
	if b.Cell(13) != 0 {
		t.Fatal("Board cell 13 is note empty")
	}
	for i := byte(0); i < 6; i++ {
		if b.Cell(i) != 4 {
			t.Fatalf("Board cell", i, "is not 4")
		}
	}
}

func TestLegalMove(t *testing.T) {
	b := MakeBoard(PlayerOne)
	b2 := MakeBoard(PlayerTwo)
	for i := byte(0); i < 6; i++ {
		if !b.LegalMove(i) {
			t.Fatalf("Player 1 cell", i, "should be legal")
		}
		if b2.LegalMove(i) {
			t.Fatalf("Player 2 cell", i, "should be illegal")
		}
	}
	for i := byte(7); i < 13; i++ {
		if b.LegalMove(i) {
			t.Fatalf("Player 1 cell", i, "should be illegal")
		}
		if !b2.LegalMove(i) {
			t.Fatalf("Player 2 cell", i, "should be legal")
		}
	}
	if b.LegalMove(6) || b.LegalMove(13) {
		t.Fatalf("Moves in kalah goals should not be legal")
	}
}

func assertMove(t *testing.T, b Board, cell byte, expectNumCaptured byte,
	expectOk bool) {
	numCaptured, ok := b.Move(cell)
	if numCaptured != expectNumCaptured || ok != expectOk {
		t.Error("Move:", cell, "failed\n\tGot: (", numCaptured, ",",
			ok, ")\n\tExpected: (", expectNumCaptured, ",", expectOk, ")")
	}
}

func assertBoard(t *testing.T, b Board, cells []byte) {
	error := false
	for i, cell := range cells {
		if cell != b.Cell(byte(i)) {
			error = true
			t.Error("Board mismatch: cell", i, ":", cell, "!=",
				b.Cell(byte(i)))
		}
	}
	if error {
		t.Error(b)
	}
}

func assertWhoseMove(t *testing.T, b Board, expectedPlayer playerId) {
	if b.WhoseTurn() != expectedPlayer {
		t.Fatalf("Current player mismatch: got %v, expected: %v", b.WhoseTurn(),
			expectedPlayer)
	}
}

//TestMove tests the board.Move function.
func TestMove(t *testing.T) {
	b := MakeBoard(PlayerOne)
	assertWhoseMove(t, b, PlayerOne)
	assertBoard(t, b, []byte{4, 4, 4, 4, 4, 4, 0,
							 4, 4, 4, 4, 4, 4, 0})
	assertMove(t, b, 0, 0, true)
	assertWhoseMove(t, b, PlayerTwo)
	assertBoard(t, b, []byte{0, 5, 5, 5, 5, 4, 0,
							4, 4, 4, 4, 4, 4, 0})
	assertMove(t, b, 3, 0, false)
	//assertWhoseMove(t, b, PlayerTwo)
	assertBoard(t, b, []byte{0, 5, 5, 5, 5, 4, 0,
							4, 4, 4, 4, 4, 4, 0})
	assertMove(t, b, 9, 0, true)
	assertBoard(t, b, []byte{0, 5, 5, 5, 5, 4, 0,
							4, 4, 0, 5, 5, 5, 1})
	assertMove(t, b, 12, 0, true)
	assertBoard(t, b, []byte{1, 6, 6, 6, 5, 4, 0,
							4, 4, 0, 5, 5, 0, 2})
	cells := []byte{0, 1, 0, 0, 0, 0, 0,
					 4, 4, 4, 5, 4, 4, 2}
	b = MakeBoardFromCells(PlayerOne, cells)
	assertMove(t, b, 1, 6, true)
	assertBoard(t, b, []byte{0, 0, 0, 0, 0, 0, 6,
							 4, 4, 4, 0, 4, 4, 2})
	assertMove(t, b, 10, 0, false)
	cells = []byte{0, 0, 0, 4, 0, 0, 0,
					 0, 1, 0, 0, 0, 0, 0}
	b = MakeBoardFromCells(PlayerTwo, cells)
	assertMove(t, b, 8, 5, true)
	assertBoard(t, b, []byte{0, 0, 0, 0, 0, 0, 0,
					 0, 0, 0, 0, 0, 0, 5})
}
