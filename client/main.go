package main

import (
	"context"
	"log"

	"teko_game/pkg/netcode"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil 
		log.Fatal(err)
	}

	c := netcode.NewGameComClient(conn)

	c.SetGroupPseudonym(context.Background())

}
