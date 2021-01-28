package middleware

import (
	"context"

	"demo.gateway.test/grpc/auth"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ServerTracing(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	parentSpanContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, auth.MetadataTextMap{md})
	spanOpts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: string(ext.Component), Value: "grpc"},
		ext.SpanKindRPCServer,
		ext.RPCServerOption(parentSpanContext),
	}
	span := opentracing.GlobalTracer().StartSpan(info.FullMethod, spanOpts...)

	defer span.Finish()

	opentracing.ContextWithSpan(ctx, span)
	return handler(ctx, req)
}
