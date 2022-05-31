package main

import (
	"context"
	"log"
	"net"

	"github.com/pandalzy/byte-douyin/pb"
	"google.golang.org/grpc"
)

type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

func (p *UserServiceImpl) Login(ctx context.Context, args *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	reply := &pb.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  "success",
		UserID:     1,
		Token:      "token",
	}
	return reply, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceImpl{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
