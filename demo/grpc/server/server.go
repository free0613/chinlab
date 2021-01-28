package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"demo.gateway.test/grpc/middleware"
	pb "demo.gateway.test/grpc/pbfile"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = "8818"

type SearchUser struct{}
type server struct{}

type Tag struct {
	Id    int64
	Name  string
	State uint
}
type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:",omitempty"`
}

func (s server) GetTagList(ctx context.Context, request *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	tag := &pb.Tag{Id: 12, Name: "dasfsa", State: 1000000}
	a := make([]*pb.Tag, 5)
	a = append(a, tag)
	return &pb.GetTagListReply{
		List: a,
		Pager: &pb.Pager{
			Page:    10,
			PageNum: 2,
		},
	}, nil
}

func (user *SearchUser) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResult, error) {
	fmt.Println("grpc server")
	return &pb.SearchResult{
		Userinfo: "server ok",
	}, nil
}

func main() {
	/*	s := grpc.NewServer()
		pb.RegisterSearchUserServer(s, &SearchUser{})
		listen, _ := net.Listen("tcp", ":"+port)
		s.Serve(listen)*/
	err := RunServer(port)
	if err != nil {
		log.Fatal("run err")
	}

}

func runHttpServer(port string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	return mux
}

func runGrpcServer(port string) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			HelloInterceptor,
			middleware.ServerTracing,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server{})
	//reflection.Register(s)
	return s
}

func runGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	serveMux := runtime.NewServeMux()

	options := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), serveMux, endpoint, options)
	return serveMux
}

func RunServer(port string) error {
	httpServer := runHttpServer(port)
	grpcServer := runGrpcServer(port)
	gateway := runGatewayServer()
	runtime.HTTPError = gatewayError
	httpServer.Handle("/", gateway)
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcServer, httpServer))

}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})

}

func gatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httperr := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httperr.Code = v.Code
			httperr.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httperr)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("拦截器")
	resp, err := handler(ctx, req)
	log.Println("bye,bye")
	return resp, err
}
