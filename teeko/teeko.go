package teeko

import (
	"math/rand"
	"time"
)

type TTEntry struct {
	depth    int
	score    float64
	flag     int
	bestMove Move
}

const (
	Exact      = 0
	LowerBound = 1
	UpperBound = 2
	Empty      = 0
	Player1    = 1
	Player2    = 2
	MaxDepth   = 5
)

type Move struct {
	FromX, FromY int // Original position of the piece (set to 0 for placing new pieces)
	ToX, ToY     int // New position to move to
}

type Teeko struct {
	Board         [25]int32
	CurrentPlayer int32
	Hash          uint64
}

func NewTeeko(board [25]int32, player int32) *Teeko {
	return &Teeko{
		Board:         board,
		CurrentPlayer: player,
	}
}

func (game *Teeko) MakeMove(move Move) {
	game.Board[(move.FromY*5 + move.FromX)] = 0
	game.Board[move.ToY*5+move.ToX] = game.CurrentPlayer
}

func (game *Teeko) IsGameOver() bool {
	// Rows
	for row := 0; row < 5; row++ {
		if (game.Board[row*5] == game.Board[row*5+1] && game.Board[row*5+1] == game.Board[row*5+2] && game.Board[row*5+2] == game.Board[row*5+3] && game.Board[row*5] != 0) ||
			(game.Board[row*5+1] == game.Board[row*5+2] && game.Board[row*5+2] == game.Board[row*5+3] && game.Board[row*5+3] == game.Board[row*5+4] && game.Board[row*5+1] != 0) {
			return true
		}
	}

	// Cols
	for col := 0; col < 5; col++ {
		if (game.Board[col] == game.Board[col+5] && game.Board[col+5] == game.Board[col+10] && game.Board[col+10] == game.Board[col+15] && game.Board[col] != 0) ||
			(game.Board[col+5] == game.Board[col+10] && game.Board[col+10] == game.Board[col+15] && game.Board[col+15] == game.Board[col+20] && game.Board[col+5] != 0) {
			return true
		}
	}

	// Check squares (2x2) eto ne sdelano
	for i := 0; i < 4; i++ {
		if game.Board[i] == game.Board[i+1] && game.Board[i+1] == game.Board[i+5] && game.Board[i+5] == game.Board[i+6] && game.Board[i] != 0 {
			return true
		}
	}

	for i := 0; i < 15; i++ {
		if (i%5 != 4) && (i^5 != 0) {
			if game.Board[i] == game.Board[i+4] && game.Board[i+4] == game.Board[i+6] && game.Board[i+6] == game.Board[i+10] && game.Board[i] != 0 {
				return true
			}
		}
	}

	if (game.Board[1] != 0 && game.Board[1] == 1 && game.Board[9] == 1 && game.Board[15] == 1 && game.Board[23] == 1) || (game.Board[1] != 0 && game.Board[1] == 2 && game.Board[9] == 2 && game.Board[15] == 2 && game.Board[23] == 2) {
		return true
	}
	if (game.Board[3] != 0 && game.Board[3] == 1 && game.Board[5] == 1 && game.Board[19] == 1 && game.Board[21] == 1) || (game.Board[3] != 0 && game.Board[3] == 2 && game.Board[5] == 2 && game.Board[19] == 2 && game.Board[21] == 2) {
		return true
	}

	if (game.Board[2] != 0 && game.Board[2] == 1 && game.Board[10] == 1 && game.Board[14] == 1 && game.Board[22] == 1) || (game.Board[2] != 0 && game.Board[2] == 2 && game.Board[10] == 2 && game.Board[14] == 2 && game.Board[22] == 2) {
		return true
	}

	return false
}

func (game *Teeko) Evaluate() float32 {
	return 1
}

func (game *Teeko) GeneratePossibleMoves() []Move {
	var moves []Move

	// Generating moves for placing a new piece on the board
	if game.isInitialPhase() {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if game.Board[y*5+x] == 0 { // Check if the cell is empty
					moves = append(moves, Move{0, 0, x, y})
				}
			}
		}
	} else {
		// Generating moves for moving an existing piece
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if game.Board[y*5+x] == game.CurrentPlayer {
					for dy := -1; dy <= 1; dy++ {
						for dx := -1; dx <= 1; dx++ {
							newX, newY := x+dx, y+dy
							if newX >= 0 && newX < 5 && newY >= 0 && newY < 5 && game.Board[newY*5+newX] == 0 {
								moves = append(moves, Move{x, y, newX, newY})
							}
						}
					}
				}
			}
		}
	}
	return moves
}

func (game *Teeko) isInitialPhase() bool {
	n := 0
	for i := 0; i < 25; i++ {
		if game.Board[i] != 0 {
			n++
		}
	}
	if n < 8 {
		return true
	} else {
		return false // Placeholder return statement
	}
}

func (game *Teeko) computeHash() {
	var hash uint64 = 0
	for i, piece := range game.Board {
		if piece != Empty {
			hash ^= zobristTable[i][piece]
		}
	}
	game.Hash = hash
}

// Zobrist table
var zobristTable [25][3]uint64
var transpositionTable = make(map[uint64]TTEntry)

func (game *Teeko) initZobristTable() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		for j := 0; j < 3; j++ {
			zobristTable[i][j] = rand.Uint64()
		}
	}
}
