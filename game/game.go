package game

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"teko_game/pkg/netcode"
	"teko_game/pkg/tko"
	PVS "teko_game/principal_variation_search"
	"teko_game/teeko"

	"google.golang.org/grpc"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Magenta = "\033[35m"
)

type Bot struct {
	MatchToken        string
	UserToken         string
	Wins              int
	Loses             int
	GameServerAddress string
	Client            netcode.GameComClient
	MatchIDPacket     netcode.MatchIDPacket
	AuthPacket        netcode.AuthPacket
	BeginningPlayer   bool
}

func NewBot(userToken string, Channel *grpc.ClientConn) *Bot {
	return &Bot{
		UserToken:  userToken,
		Client:     netcode.NewGameComClient(Channel),
		AuthPacket: netcode.AuthPacket{MatrNumber: "11824691", Secret: "Iqzwersolonew15_"},
		// MatchIDPacket: netcode.MatchIDPacket{UserToken: userToken, MatchToken: "tko#b5e2375e3baf105f9feae729e61bcd1a706df2e6ead44ba8"},
	}
}

func (bot *Bot) GetUserTokenBot() string {
	response, err := bot.Client.GetUserToken(context.Background(), &bot.AuthPacket)
	if err != nil {
		log.Fatal(err)
	}

	return response.UserToken
}

func (bot *Bot) NewMatch() {
	params := netcode.MatchRequest_TkoGameParameters{TkoGameParameters: &tko.GameParameter{}}
	request := netcode.MatchRequest{
		UserToken:                bot.UserToken,
		GameToken:                "tko",
		TimeoutSuggestionSeconds: 3600,
		GameParameters:           &params,
	}

	response, err := bot.Client.NewMatch(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NewMatch:", response.MatchToken)
	fmt.Println("First Player?:", response.BeginningPlayer)

	bot.MatchToken = response.MatchToken
	bot.MatchIDPacket = netcode.MatchIDPacket{UserToken: bot.UserToken, MatchToken: response.MatchToken}
	bot.BeginningPlayer = response.BeginningPlayer
}

func (bot *Bot) ShowElo() {
	response, err := bot.Client.GetElo(context.Background(), &netcode.IDPacket{UserToken: bot.UserToken})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.GroupElo)
}

func (bot *Bot) OpponentInfo() error {
	response, err := bot.Client.GetOpponentInfo(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (bot *Bot) GetGameStateArray() []int32 {
	response, err := bot.Client.GetGameState(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		return nil
	}

	int32Values := response.GetTkoGameState().GetBoard()

	for i := 0; i < 5; i++ {
		fmt.Print(to2DArray([25]int32(int32Values))[i])
		fmt.Print("\n")
	}
	return int32Values
}

func (bot *Bot) GetGameStatusCode() netcode.GameStatus {
	response, err := bot.Client.GetGameState(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		log.Fatal(err)
		return 7
	}
	return response.GameStatus
}

func (bot *Bot) AbortMatch() {
	response, err := bot.Client.AbortMatch(context.Background(), &bot.MatchIDPacket)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func (bot *Bot) SubmitTurn(x1 uint32, y1 uint32, x2 uint32, y2 uint32) {
	tko_game_turn := tko.GameTurn{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}

	game_turn := netcode.TurnRequest_TkoGameTurn{
		TkoGameTurn: &tko_game_turn,
	}

	request := netcode.TurnRequest{
		MatchId:  &bot.MatchIDPacket,
		GameTurn: &game_turn,
	}

	_, err := bot.Client.SubmitTurn(context.Background(), &request)

	if err != nil {
		log.Fatal(err)
	}
}

func (bot *Bot) AutoPlay() error {
	bot.NewMatch()
	var player int32
	if bot.BeginningPlayer == true {
		player = 1
	} else {
		player = 2
	}
	teeko_game := teeko.NewTeeko([25]int32(bot.GetGameStateArray()), player)
	fmt.Println("Joining the match...")
	time.Sleep(3 * time.Second)

	bot.OpponentInfo()
	bot.ShowElo()

	opponent_wait := 0
	turn := 0

	for {
		codeFromServer := bot.GetGameStatusCode()

		switch codeFromServer {
		case 0:
			// TODOBEGIN: 米莎，这里还有工作要做 (+15 或 -100 学分)

			if turn < 4 {
				teeko_game.Board = [25]int32(bot.GetGameStateArray())
				value, move := PVS.MiniMaxAlphaBeta(teeko_game, 7, math.MinInt64, math.MaxInt64, true)
				fmt.Printf("Bot played: with %d from %d,%d to %d,%d\n", value, move.FromX, move.FromY, move.ToX, move.ToY)
				bot.SubmitTurn(0, 0, uint32(move.ToX), uint32(move.ToY))
			} else {
				teeko_game.Board = [25]int32(bot.GetGameStateArray())
				value, move := PVS.MiniMaxAlphaBeta(teeko_game, 7, math.MinInt64, math.MaxInt64, true)
				fmt.Printf("Bot played: with %d from %d,%d to %d,%d\n", value, move.FromX, move.FromY, move.ToX, move.ToY)
				bot.SubmitTurn(uint32(move.FromX), uint32(move.FromY), uint32(move.ToX), uint32(move.ToY))
			}
			// TODOEND
			turn++

			opponent_wait = 0
		case 1:
			if opponent_wait == 0 {
				fmt.Println("Wait for opponent to make a move!")
				opponent_wait = 1
			}
		case 3:
			fmt.Println(Green + "MATCH OVER! We won!" + Reset)
			return nil
		case 4:
			fmt.Println(Red + "Mission faild. We'll get 'em next time!" + Reset)
			return nil
		case 5:
			fmt.Println(Magenta + "Draw!" + Reset)
			return nil
		case 6:
			fmt.Println("Match not started!")
		case 7:
			fmt.Println("Match aborted! Disconnecting...")
			time.Sleep(1 * time.Second)
			return nil
		default:
			fmt.Println("Unknown Code!")
			return nil
		}
		time.Sleep(5 * time.Second)
	}
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
