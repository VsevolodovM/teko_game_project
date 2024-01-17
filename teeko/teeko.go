package teeko

import (
	"math/rand"
	"time"
)

const (
	Exact    = 0
	Empty    = 0
	MaxDepth = 6
)

type Move struct {
	FromX, FromY int // Original position of the piece (set to 0 for placing new pieces)
	ToX, ToY     int // New position to move to
}

type Teeko struct {
	Board         [25]int32
	CurrentPlayer int32
	Hash          uint64
	ZobristTable  [25][3]uint64
}

func NewTeeko(board [25]int32, player int32) *Teeko {
	return &Teeko{
		Board:         board,
		CurrentPlayer: player,
	}
}

//	func (game *Teeko) MakeMove(move Move) {
//		if move.FromY != -1 || move.FromX != -1 {
//			game.Board[move.FromY*5+move.FromX] = 0
//			game.Board[move.ToY*5+move.ToX] = game.CurrentPlayer
//		} else {
//			game.Board[move.ToY*5+move.ToX] = game.CurrentPlayer
//		}
//		game.CurrentPlayer = 3 - game.CurrentPlayer
//		game.ComputeHash()
//	}
func (game *Teeko) MakeMove(move Move) {
	// Check if it's a placement or a shift
	if move.FromX == -1 && move.FromY == -1 {
		// It's a new piece placement
		game.Board[move.ToY*5+move.ToX] = game.CurrentPlayer
	} else {
		// It's moving an existing piece
		game.Board[move.FromY*5+move.FromX] = Empty
		game.Board[move.ToY*5+move.ToX] = game.CurrentPlayer
	}

	// Update other necessary game state information
	// For example, switch players
	game.CurrentPlayer = 3 - game.CurrentPlayer
}

// func (game *Teeko) UndoMove(move Move) {
// 	game.CurrentPlayer = 3 - game.CurrentPlayer
// 	// Reverse the move on the board
// 	if move.FromX == -1 && move.FromY == -1 {
// 		// If it was a new piece placement, just remove the piece
// 		game.Board[move.ToY*5+move.ToX] = Empty
// 	} else {
// 		// If it was a move, swap back the pieces
// 		game.Board[move.FromY*5+move.FromX] = game.CurrentPlayer
// 		game.Board[move.ToY*5+move.ToX] = Empty
// 	}

// 	// Update the Zobrist hash
// 	// For removing a piece
// 	game.Hash ^= game.ZobristTable[move.ToY*5+move.ToX][game.CurrentPlayer]
// 	if move.FromX != -1 && move.FromY != -1 {
// 		// For putting back the original piece
// 		game.Hash ^= game.ZobristTable[move.FromY*5+move.FromX][game.CurrentPlayer]
// 	}

// 	// Switch the player back
// 	game.ComputeHash()
// }

func (game *Teeko) UndoMove(move Move) {
	// Revert the move
	game.CurrentPlayer = 3 - game.CurrentPlayer
	if move.FromX == -1 && move.FromY == -1 {
		// Revert a new piece placement
		game.Board[move.ToY*5+move.ToX] = Empty
	} else {
		// Revert moving an existing piece
		game.Board[move.FromY*5+move.FromX] = game.CurrentPlayer
		game.Board[move.ToY*5+move.ToX] = Empty
	}

	game.Hash ^= game.ZobristTable[move.ToY*5+move.ToX][game.CurrentPlayer]
	if move.FromX != -1 && move.FromY != -1 {
		// For putting back the original piece
		game.Hash ^= game.ZobristTable[move.FromY*5+move.FromX][game.CurrentPlayer]
	}

	// Revert other game state changes
	game.ComputeHash()
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

	// Check squares (2x2)
	for i := 0; i < 19; i++ {
		if i%5 != 4 {
			if game.Board[i] == game.Board[i+1] && game.Board[i+1] == game.Board[i+5] && game.Board[i+5] == game.Board[i+6] && game.Board[i] != 0 {
				return true
			}
		}
	}

	// Check squares
	for i := 0; i < 15; i++ {
		if (i%5 != 4) && (i^5 != 0) {
			if game.Board[i] == game.Board[i+4] && game.Board[i+4] == game.Board[i+6] && game.Board[i+6] == game.Board[i+10] && game.Board[i] != 0 {
				return true
			}
		}
	}

	// 8 times (normal)
	for i := 2; i < 4; i++ {
		if game.Board[i] == game.Board[i+3] && game.Board[i+3] == game.Board[i+11] && game.Board[i+11] == game.Board[i+14] && game.Board[i] != 0 {
			return true
		}
		next_row := i + 5
		if game.Board[next_row] == game.Board[next_row+3] && game.Board[next_row+3] == game.Board[next_row+11] && game.Board[next_row+11] == game.Board[next_row+14] && game.Board[next_row] != 0 {
			return true
		}
	}

	// 8 times (mirrored)
	for i := 1; i < 3; i++ {
		if game.Board[i] == game.Board[i+7] && game.Board[i+7] == game.Board[i+9] && game.Board[i+9] == game.Board[i+16] && game.Board[i] != 0 {
			return true
		}
		next_row := i + 5
		if game.Board[next_row] == game.Board[next_row+7] && game.Board[next_row+7] == game.Board[next_row+9] && game.Board[next_row+9] == game.Board[next_row+16] && game.Board[next_row] != 0 {
			return true
		}
	}

	// 2 times
	if (game.Board[1] != 0 && game.Board[1] == 1 && game.Board[9] == 1 && game.Board[15] == 1 && game.Board[23] == 1) || (game.Board[1] != 0 && game.Board[1] == 2 && game.Board[9] == 2 && game.Board[15] == 2 && game.Board[23] == 2) {
		return true
	}
	if (game.Board[3] != 0 && game.Board[3] == 1 && game.Board[5] == 1 && game.Board[19] == 1 && game.Board[21] == 1) || (game.Board[3] != 0 && game.Board[3] == 2 && game.Board[5] == 2 && game.Board[19] == 2 && game.Board[21] == 2) {
		return true
	}

	// 1 time
	if (game.Board[2] != 0 && game.Board[2] == 1 && game.Board[10] == 1 && game.Board[14] == 1 && game.Board[22] == 1) || (game.Board[2] != 0 && game.Board[2] == 2 && game.Board[10] == 2 && game.Board[14] == 2 && game.Board[22] == 2) {
		return true
	}

	return false
}

func (game *Teeko) Evaluate() float32 {
	if game.IsGameOver() {
		return 100
	}
	return float32(len(game.GeneratePossibleMoves()))
}

func (game *Teeko) GeneratePossibleMoves() []Move {
	var moves []Move

	// Generating moves for placing a new piece on the board
	if game.IsInitialPhase() {
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
							// if dy == 0 && dx == 0 {
							// 	continue // Skip the current square
							// }
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
func (game *Teeko) GeneratePossibleMovesOpponent() []Move {
	var moves []Move

	// Generating moves for placing a new piece on the board
	if game.IsInitialPhase() {
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
				if game.Board[y*5+x] == 3-game.CurrentPlayer {
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

func (game *Teeko) IsInitialPhase() bool {
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

func (game *Teeko) ComputeHash() {
	var hash uint64 = 0
	for i, piece := range game.Board {
		if piece != Empty {
			hash ^= game.ZobristTable[i][piece]
		}
	}
	game.Hash = hash
}

func (game *Teeko) InitZobristTable() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		for j := 0; j < 3; j++ {
			game.ZobristTable[i][j] = rand.Uint64()
		}
	}
}

// func (game *Teeko) SwitchPlayer() {
// 	game.CurrentPlayer = 3 - game.CurrentPlayer
// }
