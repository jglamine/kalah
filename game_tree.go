package kalah

import (
		"math"
		"time"
		)

// gameState is a struct used to hold game states in nodes
// it is used internally in the gameTree graph structure
// it implements Board
type gameState struct {
	board
}

// gameTree is a graph structure which holds the game tree
type gameTree struct {
	root *gameNode
}

// gameNode - the nodes which make up a gameTree graph structure
type gameNode struct {
	state *gameState
	children []*gameNode
	move byte
}

// creates a gameTree from a Board interface
func makeGameTree(b Board) (tree *gameTree) {
	// convert Board to *state
	tree = new(gameTree)
	tree.root = new(gameNode)
	tree.root.state = boardToState(b)
	return tree
}

// boardToState creates a gameState from a Board interface
func boardToState(b Board) (s *gameState) {
	s = new(gameState)
	s.playerToMove = b.WhoseTurn()
	s.cells = make([]byte, numCells)
	for i := range s.cells {
		s.cells[i] = b.Cell(byte(i))
	}
	return s
}

func (this gameState) value(maximize bool) int64 {
	p1Score, p2Score := this.Score()
	if (this.WhoseTurn() == PlayerOne) == maximize {
		return int64(p1Score - p2Score)
	}
		return int64(p2Score - p1Score)
}

func makeGameNode(move byte, parentState *gameState) *gameNode {
	node := new(gameNode)
	node.state = new(gameState)
	node.state.cells = make([]byte, numCells)
	copy(node.state.cells, parentState.cells)
	node.state.playerToMove = parentState.playerToMove
	node.move = move
	node.state.Move(move)
	return node
}

func makeChildren(state *gameState) []*gameNode {
	if state.GameOver() {
		return make([]*gameNode, 0)
	}
	var lowerMove byte
	var upperMove byte
	if state.WhoseTurn() == PlayerOne {
		lowerMove = byte(0)
		upperMove = kalahOne - 1
	} else {
		lowerMove = kalahOne + 1
		upperMove = kalahTwo - 1
	}
	children := make([]*gameNode, 0, upperMove - lowerMove)
	for move := lowerMove; move <= upperMove; move++ {
		if state.LegalMove(move) {
			children = append(children, makeGameNode(move, state))
		}
	}
	return children
}

func (this gameTree) BestMove(timeLimit time.Duration, depthLimit int) (move byte, depth int) {
	move = kalahOne  // initialize to invalid move
	prevMove := this.search(0, time.Now())
	deadline := time.Now().Add(timeLimit)
	for depth = 1; time.Now().Before(deadline) && depth <= depthLimit; depth++ {
		prevMove = move
		move = this.search(depth, deadline)
	}
	if depth == depthLimit+1 && time.Now().Before(deadline) {
		return move, depth-1
	}
	return prevMove, depth-1
}

func (this gameTree) search(depthLimit int, deadline time.Time) (move byte) {
	if depthLimit < 1 || time.Now().After(deadline) {
		// return the first valid move
		for move = byte(0); move < numCells; move++ {
			if this.root.state.LegalMove(move) {
				return move
			}
		}
		// default to an invalid move
		return kalahOne
	}
	var alpha int64 = math.MinInt64
	var beta int64 = math.MaxInt64
	if this.root.children == nil {
		this.root.children = makeChildren(this.root.state)
	}
	if len(this.root.children) == 0 {
		// no valid moves, return an invalid move
		return kalahOne
	}
	var result int64
	var maximize bool
	for _, child := range this.root.children {
		maximize = this.root.state.WhoseTurn() == child.state.WhoseTurn()
		result = minimax(child, alpha, beta, depthLimit - 1, maximize,
			deadline)
		if result > alpha {
			alpha = result
			move = child.move
		}
		// TODO: try moving inside above if statement
		if alpha >= beta {
			return move
		}
	}
	return move
}

// minimax does minimax search with alpha beta pruning on a gameNode
// it returns the value of the given node
// depth limited breadth first search is used
// if maximize is true, maximize the score
func minimax(node *gameNode, alpha int64, beta int64, depth int, maximize bool,
	deadline time.Time) (value int64) {
	if depth == 0 || time.Now().After(deadline) {
		return node.state.value(maximize)
	}
	if node.children == nil {
		node.children = makeChildren(node.state)
	}
	if len(node.children) == 0 {
		return node.state.value(maximize)
	}
	var result int64
	if maximize {
		for _, child := range node.children {
			maximize = node.state.WhoseTurn() == child.state.WhoseTurn()
			result = minimax(child, alpha, beta, depth-1, maximize, deadline)
			if result > alpha {
				alpha = result
			}
			if alpha >= beta {
				return alpha
			}
		}
		return alpha
	} else {
		for _, child := range node.children {
			maximize = node.state.WhoseTurn() != child.state.WhoseTurn()
			result = minimax(child, alpha, beta, depth-1, maximize, deadline)
			if result < beta {
				beta = result
			}
			if beta <= alpha {
				return beta
			}
		}
		return beta
	}
}
