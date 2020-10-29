package main

import (
	"context"
	pb "demo.gateway.test/grpc/pbfile"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const port = "8818"

type SearchUser struct{}

func (user *SearchUser) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResult, error) {
	fmt.Println("grpc server")
	return &pb.SearchResult{
		Userinfo: "server ok",
	}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterSearchUserServer(s, &SearchUser{})
	listen, _ := net.Listen("tcp", ":"+port)
	s.Serve(listen)
}
