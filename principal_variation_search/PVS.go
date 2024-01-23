package PVS

import (
	"math"
	"sort"
	"teko_game/teeko"
)

// func PrincipalVariationSearch(game *teeko.Teeko, depth int, alpha int, beta int, maximizingPlayer bool) (int, teeko.Move) {

// 	if GameOver, _ := game.IsGameOver(); depth == 0 || GameOver {
// 		return game.Evaluate(game.CurrentPlayer), teeko.Move{}
// 	}

// 	var bestMove teeko.Move
// 	isFirstMove := true

// 	possibleMoves := game.GeneratePossibleMoves()

// 	for _, move := range possibleMoves {
// 		game.MakeMove(move)

// 		var score int
// 		if isFirstMove {
// 			score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer)
// 			isFirstMove = false
// 		} else {
// 			// Null window search
// 			score, _ = PrincipalVariationSearch(game, depth-1, -alpha-1, -alpha, !maximizingPlayer)
// 			if alpha < score && score < beta {
// 				// Full re-search
// 				score, _ = PrincipalVariationSearch(game, depth-1, alpha, beta, !maximizingPlayer)
// 			}
// 		}

// 		game.UndoMove(move)

// 		if maximizingPlayer {
// 			if score > alpha {
// 				alpha = score
// 				bestMove = move
// 			}
// 		} else {
// 			if score < beta {
// 				beta = score
// 				bestMove = move
// 			}
// 		}

// 		if alpha >= beta {
// 			break
// 		}
// 	}

// 	if maximizingPlayer {
// 		return alpha, bestMove
// 	} else {
// 		return beta, bestMove
// 	}
// }

// func BestMovePV(game *teeko.Teeko, transpositionTable map[uint64]TTEntry, player int32) teeko.Move {
// 	current_pl := player
// 	_, move := PrincipalVariationSearch(game, teeko.MaxDepth, math.MinInt64, math.MaxInt64, game.CurrentPlayer == current_pl, transpositionTable)
// 	return move
// }

// func PVS(game *teeko.Teeko, depth int, alpha, beta int, maximizingPlayer bool) (int, teeko.Move) {

// if GameOver, _ := game.IsGameOver(); depth == 0 || GameOver {
// 	//print(int(game.Evaluate()))
// 	return game.Evaluate(game.CurrentPlayer), teeko.Move{} // Assuming evaluate() returns the heuristic value
// }

// isPVNode := false
// var bestMove teeko.Move

// possibleMoves := game.GeneratePossibleMoves()
// for _, move := range possibleMoves {
// 	game.MakeMove(move)
// 	var score int
// 	if isPVNode {
// 		score, _ = PVS(game, depth-1, -alpha-1, -alpha, !maximizingPlayer)
// 		score = -1 * score
// 		if alpha < score && score < beta {
// 			score, _ = PVS(game, depth-1, -beta, -score, !maximizingPlayer)
// 			score = -1 * score
// 		}
// 	} else {
// 		score, _ = PVS(game, depth-1, -beta, -alpha, !maximizingPlayer)
// 		score = -1 * score
// 	}

// 	game.UndoMove(move)

// 	if score > alpha {
// 		alpha = score
// 		bestMove = move
// 	}

// 	if alpha >= beta {
// 		break
// 	}

// 	isPVNode = true
// }

// return alpha, bestMove
// }

// func MiniMax(game *teeko.Teeko, depth int, maximizingPlayer bool) (int, teeko.Move) {
// 	if depth == 0 || game.IsGameOver() {

// 		//game.CurrentPlayer = 3 - game.CurrentPlayer
// 		return game.Evaluate(), teeko.Move{} // Assuming evaluate() returns the heuristic value
// 	}

// 	var bestMove teeko.Move
// 	if maximizingPlayer {
// 		maxEval := math.MinInt32
// 		for _, move := range game.GeneratePossibleMoves() {
// 			game.MakeMove(move)
// 			eval, _ := MiniMax(game, depth-1, false)
// 			game.UndoMove(move)

// 			if eval > maxEval {
// 				maxEval = eval
// 				bestMove = move
// 			}
// 		}
// 		return maxEval, bestMove
// 	} else {
// 		minEval := math.MaxInt32
// 		test := game.Board
// 		for _, move := range game.GeneratePossibleMoves() {
// 			game.MakeMove(move)
// 			eval, _ := MiniMax(game, depth-1, true)
// 			game.UndoMove(move)
// 			if test != game.Board {

// 				fmt.Print("something is wrong")
// 			}

// 			if eval < minEval {
// 				minEval = eval
// 				bestMove = move
// 			}
// 		}

//			return minEval, bestMove
//		}
//	}
type ScoredMove struct {
	Move  teeko.Move
	Score int
}

func scoreMove(game *teeko.Teeko, move teeko.Move, player int32) int {
	game.MakeMove(move)
	score := game.Evaluate(player)
	game.UndoMove(move)
	return score
}
func topMoves(moves []teeko.Move, game *teeko.Teeko, n int, maximizingPlayer bool) []teeko.Move {
	scoredMoves := make([]ScoredMove, len(moves))
	currentPlayer := game.CurrentPlayer

	for i, move := range moves {
		scoredMoves[i] = ScoredMove{Move: move, Score: scoreMove(game, move, currentPlayer)}
	}

	// Sort to bring the top n moves to the front, considering maximizing or minimizing player
	sort.Slice(scoredMoves, func(i, j int) bool {
		if maximizingPlayer {
			return scoredMoves[i].Score > scoredMoves[j].Score // Maximizer wants higher scores first
		} else {
			return scoredMoves[i].Score < scoredMoves[j].Score // Minimizer wants lower scores first
		}
	})

	// Select and return the top n moves
	topNMoves := make([]teeko.Move, min(n, len(scoredMoves)))
	for i := 0; i < len(topNMoves); i++ {
		topNMoves[i] = scoredMoves[i].Move
	}
	return topNMoves
}

func MiniMaxAlphaBeta(game *teeko.Teeko, depth int, alpha, beta int, maximizingPlayer bool) (int, teeko.Move) {

	if GameOver, _ := game.IsGameOver(); depth == 0 || GameOver {
		return int(game.Evaluate(game.CurrentPlayer)), teeko.Move{} // Replace with your game's evaluation function
	}

	allMoves := game.GeneratePossibleMoves()
	topNMoves := topMoves(allMoves, game, 2, maximizingPlayer)
	var bestMove teeko.Move
	if maximizingPlayer {
		maxEval := math.MinInt32
		for _, move := range topNMoves {
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
