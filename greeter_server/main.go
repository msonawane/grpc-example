package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/msonawane/grpc-example/protos"
	"golang.org/x/net/context"
)

const (
	port = ":5051"
)

// server is used to helloworld.GreeterServer.
type server struct{}

// Implements helloworld.GreeterServer.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
