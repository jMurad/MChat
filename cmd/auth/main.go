package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/brianvoe/gofakeit"

	desc "github.com/jMurad/MChat/pkg/auth_v1"
)

const grpc_port = 50051

type server struct {
	desc.UnimplementedAuthV1Server
}

// Get
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.BeerName(),
				Email: gofakeit.Email(),
			},
			Role:      desc.Role(gofakeit.Uint32() % 2),
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id := gofakeit.Number(1111, 9999)

	log.Printf("Create user with id: %d\n", id)

	return &desc.CreateResponse{
		Id: int64(id),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	log.Printf("User with id {%d} is updated\n", id)

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	log.Printf("Delete user with id: %d\n", id)

	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_port))
	if err != nil {
		log.Fatalf("failde to listen: %v\n", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
