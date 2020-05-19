package start

import (
	"fmt"
	"log"
	"net"

	"github.com/atyagi9006/hello-grpc-gw/pkg/api"
	"github.com/atyagi9006/hello-grpc-gw/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Run() {
	//create a listener on tcp layer
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatal(err)
	}

	// create service hello service
	helloSvc := api.HelloGRPCService{}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("pkg/cert/server.crt", "pkg/cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor)}

	//create grpc service
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterPingServer(grpcServer, &helloSvc)

	//start grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
