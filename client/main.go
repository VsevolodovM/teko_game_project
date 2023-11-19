package main

import (
	"context"
	"log"
	"teko_game/pkg/netcode"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("gameserver.ist.tugraz.at:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := netcode.NewGameComClient(conn)
	// matr_number := ""
	// secret := ""

	autpack := &netcode.AuthPacket{MatrNumber: 12344, Secret: mgmmf}

	c.SetGroupPseudonym(context.Background(), &netcode.SetPseudonymRequest{Auth: autpack, Pseudonym: "Máo Zédōng_Team"})

}
