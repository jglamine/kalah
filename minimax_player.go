package kalah

type minimaxPlayer struct {
	player
}

func MakeMinimaxPlayer(name string) Player {
	p := minimaxPlayer{player{name}}
	return &p
}

func (this *minimaxPlayer) ChooseMove(b Board) byte {
	// TODO: keep tree between moves and resume search
	//       this can be accomplished with transposition tables
	// TODO: time limited iterative deepening
	tree := makeGameTree(b)
	move, _ := tree.BestMove(3)
	return move
}
