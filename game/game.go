package game

import (
	"context"
	"fmt"
	"log"
	"teko_game/pkg/netcode"
	"teko_game/pkg/tko"

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
}

func NewBot(userToken string, Channel *grpc.ClientConn) *Bot {
	return &Bot{
		UserToken: userToken,
		Client:    netcode.NewGameComClient(Channel),
	}
}

func (bot *Bot) newMatch() {
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

func (bot *Bot) opponentInfo() error {
	response, err := bot.Client.GetOpponentInfo(context.Background(), &bot.MatchIDPacket)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (bot *Bot) getGameState() {
	response, err := bot.Client.GetGameState(context.Background(), &bot.MatchIDPacket)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func (bot *Bot) submitTurn(x1 uint32, y1 uint32, x2 uint32, y2 uint32) {
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

func (bot *Bot) game() {
	// Implement the game logic here
}
