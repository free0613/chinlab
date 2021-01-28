package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type Auth struct {
	Appkey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	panic("implement me")
}

func (a *Auth) RequireTransportSecurity() bool {
	panic("implement me")
}

type MD map[string][]string

func main() {
	a := metadata.New(map[string]string{"go": "programming", "tour": "book"})
	b := metadata.Pairs("go", "programming", "tour", "book", "go", "eddycjy")

	ctx := context.Background()
	inctx := metadata.NewIncomingContext(ctx, a)
	outctx := metadata.NewOutgoingContext(ctx, b)

	fromin, _ := metadata.FromIncomingContext(inctx)
	fromout, _ := metadata.FromOutgoingContext(outctx)

	fmt.Println(fromin)
	fmt.Println(fromout)
}
