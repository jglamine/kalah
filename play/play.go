package main

import(
	"github.com/jglamine/kalah"
	"time"
)

func main() {
	//p1 := kalah.MakeMinimaxPlayer("Minimax depth 10", 6)
	//p1 := kalah.MakeBozoPlayer("Bozo")
	p1 := kalah.MakeHumanPlayer("Human")
	p2 := kalah.MakeMinimaxPlayer("Minimax 4 seconds", time.Duration(4)*time.Second)
	kalah.RunMatch(p1, p2, kalah.PlayerOne, true)
}
