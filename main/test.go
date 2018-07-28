package main

import (
	"context"
	"log"

	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial("127.0.0.1:1237", grpc.WithInsecure())
	if err != nil {
		log.Println("dial error", err)
		return
	}

	client := pb.NewGueryExecutorClient(grpcConn)

	instruction := pb.Instruction{}

	if _, err = client.SendInstruction(context.Background(), &instruction); err != nil {
		log.Println("send error", err)
		return
	}
	grpcConn.Close()

}
