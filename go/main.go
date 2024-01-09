// main.go
package main

import (
	"context"
	"fmt"
	pb "grpc_test/proto" // 替换为实际的 protobuf 生成文件路径
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddress := "localhost:50051"

	// 创建 gRPC 连接
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000)))
	if err != nil {
		log.Fatalf("Error connecting to the server: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := pb.NewGreeterClient(conn)

	// 调用 gRPC 服务的方法
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	fmt.Printf("Server response: %s\n", response.Message)
}
