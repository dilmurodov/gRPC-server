package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dilmurodov/grpc-course/pb"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Failed to connecting client to Server", err)
	}

	defer conn.Close()

	c := pb.NewBookShopClient(conn)

	reponse, err := c.GetBookList(context.Background(), &pb.GetBookListRequest{})

	if err != nil {
		log.Fatal("Failed on reques to server!", err.Error())
	}

	fmt.Println(reponse)
}