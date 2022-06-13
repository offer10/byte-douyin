package main

import (
	"github.com/offer10/byte-douyin/pb"
	"github.com/offer10/byte-douyin/relation-server/controller"
	"github.com/offer10/byte-douyin/relation-server/initialization"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	initialization.RegisterMySQL()
	server := grpc.NewServer()
	pb.RegisterRelationServiceServer(server, &controller.RelationServiceImpl{})

	listen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
