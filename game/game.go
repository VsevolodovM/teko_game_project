package game

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"teko_game/logic"
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
	BeginningPlayer   bool
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
	turn := 0
	opponentWait := 0

	for {
		codeFromServer := bot.GetGameStatusCode()

		switch codeFromServer {
		case 0:
			var pos int32
			game_state := bot.GetGameStateArray()

			if turn < 4 {
				switch turn {
				case 0:
					pos = 4
				case 1:
					pos = 9
				case 2:
					pos = 14
				case 3:
					pos = 19
				}

				x1, y1 := logic.OneDtotwoD(pos)
				bot.SubmitTurn(uint32(x1), uint32(y1), uint32(x1), uint32(y1))
			} else {
				if bot.BeginningPlayer {
					pos = logic.ChooseRandomPlace(game_state, 1)
				} else {
					pos = logic.ChooseRandomPlace(game_state, 2)
				}

				availableNeighbors := logic.AvailableNeighborCells(int(pos), game_state)

				if len(availableNeighbors) > 0 {
					chosenPos := availableNeighbors[rand.Intn(len(availableNeighbors))]
					x1, y1 := logic.OneDtotwoD(int32(pos))
					x2, y2 := logic.OneDtotwoD(int32(chosenPos))
					fmt.Println("X1: ", x1)
					fmt.Println("Y1: ", y1)
					fmt.Println("X2: ", x2)
					fmt.Println("Y2: ", y2)
					bot.SubmitTurn(uint32(x1), uint32(y1), uint32(x2), uint32(y2))
				} else {
					fmt.Println("No available neighboring cells.")
				}
			}

			turn++
			opponentWait = 0

		case 1:
			if opponentWait == 0 {
				fmt.Println("Wait for opponent to make a move!")
				opponentWait = 1
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

		fmt.Println("==================")
		game_array := bot.GetGameStateArray()
		for i := 0; i < len(game_array); i++ {
			fmt.Print(game_array[i], " ")
			if (i+1)%5 == 0 {
				fmt.Println()
			}
		}

		time.Sleep(3 * time.Second)
	}
}

// func (bot *Bot) GenerateTurn(game_state []uint32) (x1 uint32, y1 uint32, x2 uint32, y2 uint32) {

// }
