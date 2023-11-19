package game

import (
	"context"
	"fmt"
	"log"

	"teko_game/pkg/netcode"
	"teko_game/pkg/tko"
)

var user_token = "1234"
var match_token = ""

func NM(c netcode.GameComClient) {

	params := tko.GameParameter{}
	p := netcode.MatchRequest_TkoGameParameters{TkoGameParameters: &params}
	request := netcode.MatchRequest{UserToken: user_token, GameToken: "tko", TimeoutSuggestionSeconds: 3600, GameParameters: &p}
	response, err := c.NewMatch(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("NewMatch:", response.MatchToken)
	fmt.Print("First Player?:", response.BeginningPlayer)
	match_token = response.MatchToken
}

func game() {

}
