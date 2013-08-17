package kalah

import (
	"fmt"
	)

type humanPlayer struct {
	player
}

func MakeHumanPlayer(name string) Player {
	p := humanPlayer{player{name}}
	return &p
}

func (this *humanPlayer) ChooseMove(b Board) (move byte, depth int) {
	// TODO: read as byte instead of int
	moveInt := int(kalahOne) // initialize to an invalid move
	for !b.LegalMove(byte(moveInt)) {
		fmt.Print("Your move: ")
		_, err := fmt.Scanf("%d\n", &moveInt)
		if err != nil {
			fmt.Println(err)
		} else if !b.LegalMove(byte(moveInt)) {
			fmt.Printf("Illegal move: %d. Try again.\n", moveInt)
		}
	}
	return byte(moveInt), -1
}
