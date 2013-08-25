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
	leafReached bool
	value int64
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
		return int64(p1Score) - int64(p2Score)
	}
		return int64(p2Score) - int64(p1Score)
}

func makeGameNode(move byte, parentState *gameState) *gameNode {
	node := new(gameNode)
	node.state = new(gameState)
	node.state.cells = make([]byte, numCells)
	copy(node.state.cells, parentState.cells)
	node.state.playerToMove = parentState.playerToMove
	node.move = move
	node.leafReached = false
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
	prevMove, _ := this.search(0, time.Now()) // initialize to first valid move
	deadline := time.Now().Add(timeLimit)
	var done bool = false
	for depth = 1; !done && time.Now().Before(deadline) && ( depthLimit < 0 || depth <= depthLimit); depth++ {
		prevMove = move
		move, done = this.search(depth, deadline)
	}
	if done || (depth == depthLimit+1 && time.Now().Before(deadline)) {
		return move, depth-1
	}
	return prevMove, depth-1
}

// search runs a depth and time limited minimax search on the gameTree
// returns:
//   move (byte) - the best move
//   leafReached (bool) - true if the entire game tree was searched
func (this gameTree) search(depthLimit int, deadline time.Time) (move byte, leafReached bool) {
	if depthLimit < 1 || time.Now().After(deadline) {
		// return the first valid move
		for move = byte(0); move < numCells; move++ {
			if this.root.state.LegalMove(move) {
				return move, false
			}
		}
		// default to an invalid move
		return kalahOne, false
	}
	var alpha int64 = math.MinInt64
	var beta int64 = math.MaxInt64
	if this.root.children == nil {
		this.root.children = makeChildren(this.root.state)
	}
	if len(this.root.children) == 0 {
		// no valid moves, return an invalid move
		return kalahOne, true
	}
	var result int64
	var maximize bool
	this.root.leafReached = true
	for _, child := range this.root.children {
		maximize = this.root.state.WhoseTurn() == child.state.WhoseTurn()
		result, leafReached = minimax(child, alpha, beta, depthLimit - 1, maximize,
			deadline)
		if result > alpha {
			alpha = result
			move = child.move
		}
		if !leafReached {
			this.root.leafReached = false
		}
		// TODO: try moving inside above if statement
		if alpha >= beta {
			break
		}
	}
	return move, this.root.leafReached
}

// minimax does minimax search with alpha beta pruning on a gameNode
// depth limited breadth first search is used
// if maximize is true, maximize the score
// returns:
//    value (int64): the value of node
//    leafReached (bool): true if the branch was completely traversed
func minimax(node *gameNode, alpha int64, beta int64, depth int, maximize bool,
	deadline time.Time) (value int64, leafReached bool) {
	if node.leafReached {
		return node.value, true
	}
	if depth == 0 || time.Now().After(deadline) {
		return node.state.value(maximize), false
	}
	if node.children == nil {
		node.children = makeChildren(node.state)
	}
	if len(node.children) == 0 {
		node.value = node.state.value(maximize)
		node.leafReached = true
		return node.value, true
	}
	var result int64
	node.leafReached = true
	if maximize {
		for _, child := range node.children {
			maximize = node.state.WhoseTurn() == child.state.WhoseTurn()
			result, leafReached = minimax(child, alpha, beta, depth-1, maximize, deadline)
			if leafReached {
				child.children = nil
			} else {
				node.leafReached = false
			}
			if result > alpha {
				alpha = result
			}
			if alpha >= beta {
				break
			}
		}
		if node.leafReached {
			node.value = alpha
			return alpha, true
		}
		return alpha, false
	} else {
		for _, child := range node.children {
			maximize = node.state.WhoseTurn() != child.state.WhoseTurn()
			result, leafReached = minimax(child, alpha, beta, depth-1, maximize, deadline)
			if leafReached {
				child.children = nil
			} else {
				node.leafReached = false
			}
			if result < beta {
				beta = result
			}
			if beta <= alpha {
				break
			}
		}
		if node.leafReached {
			node.value = beta
			return beta, true
		}
		return beta, false
	}
}
