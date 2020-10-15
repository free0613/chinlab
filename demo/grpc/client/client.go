package main

import (
	"context"
	pb "demo.gateway.test/grpc/pbfile"
	"google.golang.org/grpc"
	"log"
)

const port = "8818"

var pernum map[string]*pb.SearchResult

func Search(client pb.SearchUserClient) error {
	ser := &pb.SearchResult{
		Userinfo: "",
	}
	pernum = make(map[string]*pb.SearchResult)
	pernum["123"] = ser
	resp, _ := client.Search(context.Background(), &pb.SearchRequest{Condition: "123", PersonNum: pernum})
	log.Printf("client.search resp: %s", resp.Userinfo)
	return nil
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewSearchUserClient(conn)

	_ = Search(client)
}
