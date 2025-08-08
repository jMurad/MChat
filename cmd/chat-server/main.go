package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/brianvoe/gofakeit"

	desc "github.com/jMurad/MChat/pkg/chat_v1"
)

const grpc_port = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

// Create
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id := gofakeit.Number(1111, 9999)

	log.Printf("Create chat with id: %d\n", id)

	return &desc.CreateResponse{
		Id: int64(id),
	}, nil
}

// Delete
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	log.Printf("Delete chat with id: %d\n", id)

	return &emptypb.Empty{}, nil
}

// SendMessage
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message from %s with text: '%s' is sended\n", req.GetMessage().GetFrom(), req.GetMessage().GetText())

	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_port))
	if err != nil {
		log.Fatalf("failde to listen: %v\n", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
