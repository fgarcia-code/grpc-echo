package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	api "github.com/fgarcia-code/grpc-echo/pkg/grpc/echo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 5001, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	api.UnimplementedEchoServiceServer
}

func (s *server) EchoUnary(ctx context.Context, message *api.EchoMessage) (*api.EchoMessage, error) {
	return message, nil
}

func (s *server) EchoClientStream(stream api.EchoService_EchoClientStreamServer) error {
	for {
		message, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(message.Message)
	}
	return nil
}

func (s *server) EchoServerStream(message *api.EchoMessage, stream api.EchoService_EchoServerStreamServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(message)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) EchoBidiStream(stream api.EchoService_EchoBidiStreamServer) error {
	for {
		message, err := stream.Recv()
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
		stream.Send(message)
	}
	return nil
}

func (s *server) EchoStatus(ctx context.Context, statusCode *api.StatusCode) (*api.StatusCode, error) {
	return statusCode, status.Error(codes.Code(statusCode.Code.Number()), statusCode.Code.String())
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	api.RegisterEchoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
