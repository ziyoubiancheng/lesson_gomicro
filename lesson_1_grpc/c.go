package main

import (
	"context"
	"fmt"
	pb "muke/grpc/proto"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:2333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)

	// 调用服务
	req := new(pb.UserRequest)
	req.Name = "hi"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}
