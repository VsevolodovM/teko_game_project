package game

import (
	"context"
	"fmt"
	"log"
	"teko_game/pkg/netcode"
	"teko_game/pkg/tko"
	"time"

	"google.golang.org/grpc"
)

type Bot struct {
	MatchToken        string
	Elo               int
	UserToken         string
	Wins              int
	Loses             int
	GameServerAddress string
	Client            netcode.GameComClient
	MatchIDPacket     netcode.MatchIDPacket
	AuthPacket        netcode.AuthPacket
}

func NewBot(userToken string, Channel *grpc.ClientConn) *Bot {
	return &Bot{
		UserToken:  userToken,
		Client:     netcode.NewGameComClient(Channel),
		AuthPacket: netcode.AuthPacket{MatrNumber: "11824691", Secret: "Iqzwersolonew15_"},
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
	bot.MatchIDPacket = netcode.MatchIDPacket{UserToken: bot.UserToken, MatchToken: bot.MatchToken}
}

func (bot *Bot) OpponentInfo() error {
	response, err := bot.Client.GetOpponentInfo(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (bot *Bot) GetGameState() []int32 {
	response, err := bot.Client.GetGameState(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		return nil
	}

	int32Values := response.GetTkoGameState().GetBoard()

	fmt.Println("Array:", int32Values)
	return int32Values
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

	response, err := bot.Client.SubmitTurn(context.Background(), &request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func (bot *Bot) AutoPlay() error {
	// Finding match newMatch:
	// - Waiting for the match to start. while response != OK

	bot.NewMatch()
	time.Sleep(3 * time.Second)

	bot.GetGameState()
	time.Sleep(3 * time.Second)

	// err := bot.waitMatchStarted()
	// if err != nil {
	// 	return err
	// }

	// for {
	// 	gameState, _, err := bot.getGameState()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	boardSize := len(gameState.Board)
	// 	lastIndex := uint32(boardSize - 1)

	// 	if rand.Intn(2) == 0 {
	// 		err := bot.submitTurn(lastIndex, lastIndex, lastIndex, lastIndex)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		x1 := uint32(rand.Intn(boardSize))
	// 		y1 := uint32(rand.Intn(boardSize))
	// 		x2 := uint32(rand.Intn(boardSize))
	// 		y2 := uint32(rand.Intn(boardSize))

	// 		err := bot.submitTurn(x1, y1, x2, y2)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// 	time.Sleep(2 * time.Second)

	// 	if gameState.GameOver {
	// 		break
	// 	}
	// }

	return nil
}
