package main

// "log"
// "time"
// "teko_game/game"
// "teko_game/game"
// "teko_game/pkg/tko"

import (
	//"bufio"
	"fmt"

	//"os"

	"math"
	"strconv"
	"strings"
	PVS "teko_game/principal_variation_search"
	"teko_game/teeko"
	"time"
)

//"google.golang.org/grpc"

func parseInput(input string, game *teeko.Teeko) teeko.Move {
	parts := strings.Fields(input)
	if len(parts) < 4 {
		return teeko.Move{FromX: -1, FromY: -1, ToX: -1, ToY: -1} // Invalid move format
	}
	fromX, err1 := strconv.Atoi(parts[0])
	fromY, err2 := strconv.Atoi(parts[1])
	toX, err3 := strconv.Atoi(parts[2])
	toY, err4 := strconv.Atoi(parts[3])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return teeko.Move{FromX: -1, FromY: -1, ToX: -1, ToY: -1} // Error in conversion
	}
	// Add additional validation based on game rules if necessary
	return teeko.Move{FromX: fromX, FromY: fromY, ToX: toX, ToY: toY}
}

func botOneMove(game *teeko.Teeko, tt *map[uint64]PVS.TTEntry) teeko.Move {
	// Implement the bot's strategy (e.g., using Principal Variation Search)
	// For simplicity, this example just returns a random valid move
	moves := game.GeneratePossibleMoves()
	fmt.Print(moves)
	if len(moves) > 0 {
		game.ComputeHash()
		move := PVS.BestMovePV(game, *tt, game.CurrentPlayer)
		//_, move := PVS.PVS(game, 7, math.MinInt64, math.MaxInt64, game.CurrentPlayer == 1, *tt)
		//_, move := PVS.MiniMax(game, 5, game.CurrentPlayer == 1)
		fmt.Print(move)
		return move
	}
	return teeko.Move{FromX: -1, FromY: -1} // No valid moves
}
func botTwoMove(game *teeko.Teeko, tt *map[uint64]PVS.TTEntry) teeko.Move {
	// Implement the bot's strategy (e.g., using Principal Variation Search)
	// For simplicity, this example just returns a random valid move
	moves := game.GeneratePossibleMoves()
	fmt.Print(moves)
	if len(moves) > 0 {
		game.ComputeHash()
		//move := PVS.BestMovePV(game, *tt, game.CurrentPlayer)
		//_, move := PVS.PVS(game, 6, math.MinInt64, math.MaxInt64, game.CurrentPlayer == 2, *tt)
		_, move := PVS.MiniMaxAlphaBeta(game, 7, math.MinInt64, math.MaxInt64, game.CurrentPlayer == 2)
		fmt.Print(move)
		return move
	}
	return teeko.Move{FromX: -1, FromY: -1} // No valid moves
}
func to2DArray(oneDArray [25]int32) [5][5]int32 {
	var twoDArray [5][5]int32

	for i := 0; i < 25; i++ {
		row := i / 5
		col := i % 5
		twoDArray[row][col] = oneDArray[i]
	}

	return twoDArray
}

func main() {

	// conn, err := grpc.Dial("gameserver.ist.tugraz.at:80", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// matr_number := ""
	// secret := ""
	// bot := game.NewBot("5d7d39275cd5d64cb9c5f42ff45f5489b73a42ffb65bd0b8d3b3bcb3fed1d574", conn)

	// for {
	// 	bot.AutoPlay()

	// 	time.Sleep(5 * time.Second)
	// }

	// autpack := &netcode.AuthPacket{MatrNumber: "12344", Secret: "mgmmf"}

	// c.SetGroupPseudonym(context.Background(), &netcode.SetPseudonymRequest{Auth: autpack, Pseudonym: "Máo Zédōng_Team"})
	// var a [25]int32
	// teeko1 := teeko.NewTeeko(a, 1)
	// var transpositionTable = make(map[uint64]PVS.TTEntry)
	// teeko1.InitZobristTable()
	// teeko1.ComputeHash()

	// for i := 0; i < 100; i++ {

	// }

	var a [25]int32
	game := teeko.NewTeeko(a, 1)
	var transpositionTable = make(map[uint64]PVS.TTEntry)
	var transpositionTable2 = make(map[uint64]PVS.TTEntry)
	game.InitZobristTable()
	game.ComputeHash()
	//reader := bufio.NewReader(os.Stdin)

	for !game.IsGameOver() {
		for i := 0; i < 5; i++ {
			fmt.Print(to2DArray(game.Board)[i])
			fmt.Print("\n")
		}
		fmt.Print("\n")
		if game.CurrentPlayer == 1 {
			time.Sleep(2 * time.Second)
			// fmt.Println("Your turn. Enter your move as 'fromX fromY toX toY':")

			// input, _ := reader.ReadString('\n')
			// move := parseInput(input, game)

			// game.MakeMove(move)
			start := time.Now()
			fmt.Println("Bot1's turn...")
			move := botOneMove(game, &transpositionTable)
			game.MakeMove(move)
			fmt.Printf("Bot played: from %d,%d to %d,%d\n", move.FromX, move.FromY, move.ToX, move.ToY)
			elapsed := time.Since(start)
			fmt.Print("time: ", elapsed, "\n")
		} else {
			time.Sleep(2 * time.Second)
			start := time.Now()
			fmt.Println("Bot2's turn...")
			move := botTwoMove(game, &transpositionTable2)
			game.MakeMove(move)
			fmt.Printf("Bot played: from %d,%d to %d,%d\n", move.FromX, move.FromY, move.ToX, move.ToY)
			elapsed := time.Since(start)
			fmt.Print("time: ", elapsed, "\n")
		}

	}
	for i := 0; i < 5; i++ {
		fmt.Print(to2DArray(game.Board)[i])
		fmt.Print("\n")
	}
	fmt.Println("Game over!")

}
