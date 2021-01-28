package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func defaaultContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		defaultTimeout := 60 * time.Second
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	}
	return ctx, cancel
}

func CheckParms() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Println("client begin 拦截")
		ctx, cancel := defaaultContextTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}

		log.Println("bye,bye")
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
