package main

import (
	"context"
	"fmt"
	"github.com/yyf-404/serialize_pb/pb"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()
	client := serialize.NewSerializeServiceClient(conn)
	req := serialize.MsgRequest{
		//Name:   "LiHua",
		Id:     5,
		IsUser: true,
	}
	client.SendMsg(context.Background(), &req)
}
