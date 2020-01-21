package server

import (
	"fmt"
	pb "github.com/rhizomata/go-concurrency/grpc/proto"
	"google.golang.org/grpc" 
	"log"
	"net"
	"context"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello ..
func (svr *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

// Start start server
func Start(port uint) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("Starting Server")
	_ = grpcServer.Serve(lis)
}
