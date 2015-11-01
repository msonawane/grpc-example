package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/msonawane/grpc-example/protos"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5051"
	defaultName = "world"
)

func main() {
	// Set up connection to server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to server %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.Message)
}
