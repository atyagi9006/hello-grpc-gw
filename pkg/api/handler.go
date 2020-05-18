package api

import (
	"context"
	"log"

	pb "github.com/atyagi9006/hello-grpc-gw/pkg/proto"
)

type HelloGRPCService struct {
}

//SayHello is implementing grpc-hello-world
func (svc *HelloGRPCService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Milgya greeting...bas rehn de....",
	}
	return &res, nil
}
