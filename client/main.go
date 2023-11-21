package main

import (
	"log"

	//"teko_game/game"
	"teko_game/game"

	//"teko_game/pkg/tko"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("gameserver.ist.tugraz.at:80", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// matr_number := ""
	// secret := ""
	bot := game.NewBot("5d7d39275cd5d64cb9c5f42ff45f5489b73a42ffb65bd0b8d3b3bcb3fed1d574", conn)

	bot.AutoPlay()
	// autpack := &netcode.AuthPacket{MatrNumber: "12344", Secret: "mgmmf"}

	// c.SetGroupPseudonym(context.Background(), &netcode.SetPseudonymRequest{Auth: autpack, Pseudonym: "Máo Zédōng_Team"})

}
