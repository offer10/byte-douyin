package main

import (
	"log"
	"net"

	"github.com/offer10/byte-douyin/interaction-server/controller"
	"github.com/offer10/byte-douyin/interaction-server/initialization"
	"github.com/offer10/byte-douyin/pb"
	"google.golang.org/grpc"
)

func main() {
	initialization.RegisterMySQL()
	server := grpc.NewServer()
	pb.RegisterFavoriteServiceServer(server, &controller.FavoriteServerImpl{})
	pb.RegisterCommentServiceServer(server, &controller.CommentServerImpl{})
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
