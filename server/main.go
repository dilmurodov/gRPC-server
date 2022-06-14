package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dilmurodov/grpc-course/server/bookShop"
	"google.golang.org/grpc"
	"github.com/dilmurodov/grpc-course/pb"

)

func main() {

	lis, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatal("Failed connecting server!", err)
		panic(err)
	}

	fmt.Println("Server listening on port 9000")

	grpcsServer := grpc.NewServer()

	book_srv := bookShop.Server{}

	pb.RegisterBookShopServer(grpcsServer, &book_srv)

	if err := grpcsServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}
}
