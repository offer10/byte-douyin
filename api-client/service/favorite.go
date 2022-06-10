package service

import (
	"github.com/offer10/byte-douyin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var FavoriteClient pb.FavoriteServiceClient

func FavoriteConn() {
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	log.Println(err)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	FavoriteClient = pb.NewFavoriteServiceClient(conn)
}
