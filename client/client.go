package main

import (
	"context"
	"log"

	"github.com/atyagi9006/hello-grpc-gw/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("pkg/cert/server.crt", "mycompany.com")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Setup the login/pass
	auth := Authentication{
		Login:    "amit",
		Password: "tyagi",
	}

	//get a connection by dialing G-RPC
	conn, err := grpc.Dial(":7777", grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := proto.NewPingClient(conn)
	res, err := client.SayHello(context.Background(), &proto.PingMessage{Greeting: "le greeting ...."})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server : %s \n", res.Greeting)

}
