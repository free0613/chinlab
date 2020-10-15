package global

import (
	"github.com/opentracing/opentracing-go"
	"log"
	"resk.micro/tracer"
)

var (
	Traceer opentracing.Tracer
)

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func setupTracer() error {
	jaegerTracer, closer, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	Traceer = jaegerTracer
	closer.Close()
	return nil
}
