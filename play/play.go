package main

import(
	"github.com/jglamine/kalah"
)

func main() {
	p1 := kalah.MakeHumanPlayer("Human")
	p2 := kalah.MakeMinimaxPlayer("Minimax")
	kalah.RunMatch(p1, p2, kalah.PlayerTwo, true)
}