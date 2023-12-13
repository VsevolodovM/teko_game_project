package PVS

import (
	"math"
	"teko_game/teeko"
)

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

func BestMovePV(game *teeko.Teeko, transpositionTable map[uint64]TTEntry) teeko.Move {
	_, move := PrincipalVariationSearch(game, teeko.MaxDepth, math.MinInt64, math.MaxInt64, game.CurrentPlayer == Player1, transpositionTable)
	return move
}
