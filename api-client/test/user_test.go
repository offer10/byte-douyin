package main

import (
	"context"
	"testing"
	"time"

	"github.com/pandalzy/byte-douyin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gotest.tools/v3/assert"
)

func TestLogin(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Equal(t, nil, err)

	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()

	resp, err := client.Login(ctx, &pb.UserLoginRequest{
		Username: "test",
		Password: "123456",
	})
	assert.Equal(t, nil, err)
	t.Logf("login resp: %+v", resp)
	assert.Equal(t, int32(200), resp.StatusCode)
}
