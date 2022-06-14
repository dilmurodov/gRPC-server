package bookShop

import (
	"context"
	"fmt"

	"github.com/dilmurodov/grpc-course/pb"
)

type Server struct {
	pb.UnimplementedBookShopServer
}

func (s *Server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	
	fmt.Println("Request commin ...")

	return &pb.GetBookListResponse{
		Book: []*pb.Book{
			{
				Name: "kings don't die",
				Pages: 234,
				Author: "John",
				Lan: "English",
			},
			{
				Name: "Ring of king",
				Pages: 214,
				Author: "Alex",
				Lan: "English",
			},
		},
	}, nil
}
