package kalah

import (
	"time"
)

type minimaxPlayer struct {
	player
	timeLimit time.Duration
	depthLimit int
}

func MakeMinimaxPlayer(name string, timeLimit time.Duration,
	depthLimit int) Player {
	p := minimaxPlayer{player{name}, timeLimit, depthLimit}
	return &p
}

func (this *minimaxPlayer) ChooseMove(b Board) byte {
	// TODO: keep tree between moves and resume search
	//       this can be accomplished with transposition tables
	// TODO: time limited iterative deepening
	tree := makeGameTree(b)
	move, _ := tree.BestMove(this.timeLimit, this.depthLimit)
	return move
}
