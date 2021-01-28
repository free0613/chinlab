package main

import (
	"context"
	"log"

	"demo.gateway.test/grpc/middleware"
	pb "demo.gateway.test/grpc/pbfile"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(middleware.CheckParms())))

	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(2),
				grpc_retry.WithCodes(
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
		),
	))

	// conn, _ := grpc.Dial(":"+port, opts...)
	conn, _ := grpc.DialContext(context.Background(), "127.0.0.1:8818", opts...)
	defer conn.Close()

	client := pb.NewTagServiceClient(conn)

	_ = Taglist(client)
}
