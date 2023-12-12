package PVS

import (
	"math"
	"teko_game/teeko"
)

func PrincipalVariationSearch(game teeko.Teeko, depth int, alpha, beta float32, maximizingPlayer bool, transpositionTable *map[uint64]teeko.TTEntry) (float32, teeko.Move) {
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
			score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer)
			isFirstMove = false
		} else {
			// Null window search
			score, _ = PrincipalVariationSearch(game, depth-1, -alpha-1, -alpha, !maximizingPlayer)
			if alpha < score && score < beta {
				// Full re-search
				score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer)
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

	if maximizingPlayer {
		return alpha, bestMove
	} else {
		return beta, bestMove
	}
}

func BestMovePV(game teeko.Teeko) teeko.Move {
	_, move := PrincipalVariationSearch(game, teeko.Empty, math.MinInt64, math.MaxInt64, game.CurrentPlayer == teeko.Player1)
	return move
}
