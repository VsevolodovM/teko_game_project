package game

import (
	"context"
	"fmt"
	"log"

	"teko_game/pkg/netcode"
	"teko_game/pkg/tko"
)

type Bot struct {
	MatchToken        string
	Elo               int
	UserToken         string
	Wins              int
	Loses             int
	GameServerAddress string
	SetUpChannel      string
	Client            netcode.GameComClient
}

func NewBot(userToken string) *Bot {
	return &Bot{
		UserToken: userToken,
	}
}

func (bot *Bot) newMatch() {
	params := tko.GameParameter{}
	p := netcode.MatchRequest_TkoGameParameters{TkoGameParameters: &params}
	request := netcode.MatchRequest{
		UserToken:                bot.UserToken,
		GameToken:                "tko",
		TimeoutSuggestionSeconds: 3600,
		GameParameters:           &p,
	}
	response, err := bot.Client.NewMatch(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NewMatch:", response.MatchToken)
	fmt.Println("First Player?:", response.BeginningPlayer)
	bot.MatchToken = response.MatchToken
}

func (bot *Bot) opponentInfo() error {
	request := netcode.MatchIDPacket{
		UserToken:  bot.UserToken,
		MatchToken: bot.MatchToken,
	}

	response, err := bot.Client.GetOpponentInfo(context.Background(), &request)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (bot *Bot) getGameState(client netcode.GameComClient) {

}

func (bot *Bot) game() {
	// Implement the game logic here
}
