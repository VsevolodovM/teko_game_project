package main

import (
	"fmt"
	//"log"
	//"time"

	//"teko_game/game"
	//"teko_game/game"

	//"teko_game/pkg/tko"
	"teko_game/teeko"
	//"google.golang.org/grpc"
)

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
	var a [25]int32
	b := [4]uint32{0, 0, 1, 1}
	teeko1 := teeko.NewTeeko(a, 1)
	fmt.Print(teeko1.MakeMove(a, b, 1))
}
