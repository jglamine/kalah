package kalah

import (
	"fmt"
)

// RunMatch runs a kalah match.
// firstMove specifies which player goes first (PlayerOne or PlayerTwo).
func RunMatch(playerOne Player, playerTwo Player, firstMove playerId, 
	verbose bool) (p1Score, p2Score byte) {
	if firstMove == PlayerOne {
		if verbose {
			fmt.Println("Player", playerOne, "moves first.")
		}
	} else {
		if verbose {
			fmt.Println("Player", playerTwo, "moves first.")
		}
	}
	b := MakeBoard(firstMove)
	if verbose {
		fmt.Println(b)
	}
	// main game loop
	move := byte(0)
	for !b.GameOver() {
		if verbose {
			fmt.Println("")
		}
		if b.WhoseTurn() == PlayerOne {
			move = playerOne.ChooseMove(b)
			if verbose {
				fmt.Println(playerOne, "chooses move", move)
			}
		} else {
			move = playerTwo.ChooseMove(b)
			if verbose {
				fmt.Println(playerTwo, "chooses move", move)
			}
		}
		b.Move(move)
		if verbose { fmt.Println(b) }
	}
	score1, score2 := b.Score()
	if verbose {
		if b.Winner() == PlayerOne {
			fmt.Println("Player", playerOne, "(p1) wins",
				score1, "to", score2)
		} else if b.Winner() == PlayerTwo {
			fmt.Println("Player", playerTwo, "(p2) wins",
				score2, "to", score1)
		} else {
			fmt.Println("Tie game:", score1, "to", score2)
		}
	}
	return score1, score2
}

