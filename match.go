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
	var ok bool
	var depth int
	var p1Moves int
	var p2Moves int
	var p1DepthSum int
	var p2DepthSum int
	move := byte(0)
	for !b.GameOver() {
		if verbose {
			fmt.Println("")
		}
		if b.WhoseTurn() == PlayerOne {
			move, depth = playerOne.ChooseMove(b)
			p1Moves++
			p1DepthSum += depth
			if verbose {
				fmt.Printf("%v chooses move %v", playerOne, move)
			}
		} else {
			move, depth = playerTwo.ChooseMove(b)
			p2Moves++
			p2DepthSum += depth
			if verbose {
				fmt.Printf("%v chooses move %v", playerTwo, move)
			}
		}
		if verbose {
			if depth >= 0 {
				fmt.Printf(" (depth %v)", depth)
			}
			fmt.Println("")
		}
		_, ok = b.Move(move)
		if !ok {
			fmt.Printf("Error: the move was invalid")
			// TODO: return a status code
			return 0, 0
		}

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
		if p1DepthSum >= 0 {
			fmt.Printf("%v depth average: %v\n", playerOne,
				float64(p1DepthSum) / float64(p1Moves))
		}
		if p2DepthSum >= 0 {
			fmt.Printf("%v depth average: %v\n", playerTwo,
				float64(p2DepthSum) / float64(p2Moves))
		}
	}
	return score1, score2
}

