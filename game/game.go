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
	Client            string
}

func NewBot(userToken string) *Bot {
	return &Bot{
		UserToken: userToken,
	}
}

func (b *Bot) newMatch(c netcode.GameComClient) {
	params := tko.GameParameter{}
	p := netcode.MatchRequest_TkoGameParameters{TkoGameParameters: &params}
	request := netcode.MatchRequest{
		UserToken:                b.UserToken,
		GameToken:                "tko",
		TimeoutSuggestionSeconds: 3600,
		GameParameters:           &p,
	}
	response, err := c.NewMatch(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NewMatch:", response.MatchToken)
	fmt.Println("First Player?:", response.BeginningPlayer)
	b.MatchToken = response.MatchToken
}

func (b *Bot) game() {

}
