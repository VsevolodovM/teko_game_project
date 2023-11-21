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

	fmt.Println("Array:", int32Values)
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

	response, err := bot.Client.SubmitTurn(context.Background(), &request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func (bot *Bot) AutoPlay() error {
	bot.NewMatch()
	fmt.Println("Joining the match...")
	time.Sleep(3 * time.Second)

	opponent_wait := 0
	for {
		codeFromServer := bot.GetGameStatusCode()

		switch codeFromServer {
		case 0:
			// TODO: 米莎，这里还有工作要做 (+15 或 -100 学分)
			bot.SubmitTurn(1, 2, 3, 4)
			opponent_wait = 0
		case 1:
			if opponent_wait == 0 {
				fmt.Println("Wait for opponent to make a move!")
				fmt.Println(bot.GetGameStateArray())
				opponent_wait = 1
			}
		case 3:
			fmt.Println("MATCH OVER! We won!")
			return nil
		case 4:
			fmt.Println("We lost, but keep your chin up!")
			return nil
		case 5:
			fmt.Println("Draw!")
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
