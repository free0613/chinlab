package tracer

import (
	"log"

	"github.com/opentracing/opentracing-go"
	"resk.micro/infra/base"
)

var (
	Tracer opentracing.Tracer
)

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatal("init.setupTracer err")
	}
	base.Logger().Info("log trace start success")
}

func setupTracer() error {
	jaegerTracer, closer, err := NewJaegerTracer("blog-service", "192.168.3.39:6831")
	if err != nil {
		return err
	}

	Tracer = jaegerTracer
	closer.Close()
	return nil
}
