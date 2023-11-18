package main

import (
	"context"
	"log"

	"pkg/api"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewGameComClient(conn)

	c.SetGroupPseudonym(context.Background())

}
