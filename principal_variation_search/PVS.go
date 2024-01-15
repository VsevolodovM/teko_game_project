package PVS

import (
	"fmt"
	"math"
	"teko_game/teeko"
)

// remove later
func to2DArray(oneDArray [25]int32) [5][5]int32 {
	var twoDArray [5][5]int32

	for i := 0; i < 25; i++ {
		row := i / 5
		col := i % 5
		twoDArray[row][col] = oneDArray[i]
	}

	return twoDArray
}

type TTEntry struct {
	depth    int
	score    float32
	flag     int
	bestMove teeko.Move
}

const (
	Exact      = 0
	LowerBound = 1
	UpperBound = 2
	Player1    = 1
	Player2    = 2
)

func PrincipalVariationSearch(game *teeko.Teeko, depth int, alpha, beta float32, maximizingPlayer bool, transpositionTable map[uint64]TTEntry) (float32, teeko.Move) {
	origAlpha := alpha

	if depth == 0 || game.IsGameOver() {
		return game.Evaluate(), teeko.Move{}
	}

	ttEntry, found := transpositionTable[game.Hash]
	if found && ttEntry.depth >= depth {
		if ttEntry.flag == Exact {
			return ttEntry.score, ttEntry.bestMove
		} else if ttEntry.flag == LowerBound && ttEntry.score > alpha {
			alpha = ttEntry.score
		} else if ttEntry.flag == UpperBound && ttEntry.score < beta {
			beta = ttEntry.score
		}
		if alpha >= beta {
			return ttEntry.score, ttEntry.bestMove
		}
	}

	var bestMove teeko.Move
	isFirstMove := true

	possibleMoves := game.GeneratePossibleMoves()

	for _, move := range possibleMoves {
		game.MakeMove(move)

		var score float32
		if isFirstMove {
			score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer, transpositionTable)
			isFirstMove = false
		} else {
			// Null window search
			score, _ = PrincipalVariationSearch(game, depth-1, -alpha-1, -alpha, !maximizingPlayer, transpositionTable)
			if alpha < score && score < beta {
				// Full re-search
				score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer, transpositionTable)
			}
		}

		game.UndoMove(move)

		if maximizingPlayer {
			if score > alpha {
				alpha = score
				bestMove = move
			}
		} else {
			if score < beta {
				beta = score
				bestMove = move
			}
		}

		if alpha >= beta {
			break
		}
	}
	var flag int
	if alpha <= origAlpha {
		flag = UpperBound
	} else if alpha >= beta {
		flag = LowerBound
	} else {
		flag = Exact
	}
	transpositionTable[game.Hash] = TTEntry{depth, alpha, flag, bestMove}

	if maximizingPlayer {
		return alpha, bestMove
	} else {
		return beta, bestMove
	}
}

func BestMovePV(game *teeko.Teeko, transpositionTable map[uint64]TTEntry, player int32) teeko.Move {
	current_pl := player
	_, move := PrincipalVariationSearch(game, teeko.MaxDepth, math.MinInt64, math.MaxInt64, game.CurrentPlayer == current_pl, transpositionTable)
	return move
}

func PVS(game *teeko.Teeko, depth int, alpha, beta float32, maximizingPlayer bool) (float32, teeko.Move) {

	if depth == 0 || game.IsGameOver() {
		//print(int(game.Evaluate()))
		return game.Evaluate(), teeko.Move{} // Assuming evaluate() returns the heuristic value
	}

	isPVNode := false
	var bestMove teeko.Move

	possibleMoves := game.GeneratePossibleMoves()
	for _, move := range possibleMoves {
		game.MakeMove(move)
		var score float32
		if isPVNode {
			score, _ = PVS(game, depth-1, -alpha-1, -alpha, !maximizingPlayer)
			score = -1 * score
			if alpha < score && score < beta {
				score, _ = PVS(game, depth-1, -beta, -score, !maximizingPlayer)
				score = -1 * score
			}
		} else {
			score, _ = PVS(game, depth-1, -beta, -alpha, !maximizingPlayer)
			score = -1 * score
		}

		game.UndoMove(move)

		if score > alpha {
			alpha = score
			bestMove = move
		}

		if alpha >= beta {
			break
		}

		isPVNode = true
	}

	return alpha, bestMove
}

func MiniMax(game *teeko.Teeko, depth int, maximizingPlayer bool) (int, teeko.Move) {
	if depth == 0 || game.IsGameOver() {

		//game.CurrentPlayer = 3 - game.CurrentPlayer
		return int(game.Evaluate()), teeko.Move{} // Assuming evaluate() returns the heuristic value
	}

	var bestMove teeko.Move
	if maximizingPlayer {
		maxEval := math.MinInt32
		for _, move := range game.GeneratePossibleMoves() {
			game.MakeMove(move)
			eval, _ := MiniMax(game, depth-1, false)
			game.UndoMove(move)

			if eval > maxEval {
				maxEval = eval
				bestMove = move
			}
		}
		return maxEval, bestMove
	} else {
		minEval := math.MaxInt32
		test := game.Board
		for _, move := range game.GeneratePossibleMoves() {
			game.MakeMove(move)
			eval, _ := MiniMax(game, depth-1, true)
			game.UndoMove(move)
			if test != game.Board {

				fmt.Print("something is wrong")
			}

			if eval < minEval {
				minEval = eval
				bestMove = move
			}
		}

		return minEval, bestMove
	}
}

func MiniMaxAlphaBeta(game *teeko.Teeko, depth int, alpha, beta int, maximizingPlayer bool) (int, teeko.Move) {
	if depth == 0 || game.IsGameOver() {
		return int(game.Evaluate()), teeko.Move{} // Replace with your game's evaluation function
	}

	var bestMove teeko.Move
	if maximizingPlayer {
		maxEval := math.MinInt32
		for _, move := range game.GeneratePossibleMoves() {
			game.MakeMove(move)
			eval, _ := MiniMaxAlphaBeta(game, depth-1, alpha, beta, false)
			game.UndoMove(move)

			if eval > maxEval {
				maxEval = eval
				bestMove = move
			}
			alpha = max(alpha, eval)
			if beta <= alpha {
				break // Beta cut-off
			}
		}
		return maxEval, bestMove
	} else {
		minEval := math.MaxInt32
		for _, move := range game.GeneratePossibleMoves() {
			game.MakeMove(move)
			eval, _ := MiniMaxAlphaBeta(game, depth-1, alpha, beta, true)
			game.UndoMove(move)

			if eval < minEval {
				minEval = eval
				bestMove = move
			}
			beta = min(beta, eval)
			if beta <= alpha {
				break // Alpha cut-off
			}
		}
		return minEval, bestMove
	}
}

// Helper functions for min and max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
