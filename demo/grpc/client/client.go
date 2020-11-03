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

func Taglist(client pb.TagServiceClient) error {
	list, _ := client.GetTagList(context.Background(), &pb.GetTagListRequest{
		Name:  "儿童",
		State: 1,
	})
	log.Printf("client.GetTagList tag: %v", list.List)
	return nil
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewTagServiceClient(conn)

	_ = Taglist(client)
}
