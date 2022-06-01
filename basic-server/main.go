package main

import (
	"log"
	"net"

	"github.com/offer10/byte-douyin/basic-server/controller"
	"github.com/offer10/byte-douyin/basic-server/initialization"
	"github.com/offer10/byte-douyin/pb"
	"google.golang.org/grpc"
)

func main() {
	initialization.RegisterMySQL()
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &controller.UserServerImpl{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
