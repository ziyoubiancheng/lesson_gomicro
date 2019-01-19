package ztrace

import (
	"log"

	"github.com/micro/go-micro/server"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func InitTracer(zipkinURL string, hostPort string, serviceName string) {
	collector, err := zipkin.NewHTTPCollector(zipkinURL)
	if err != nil {
		log.Fatalf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, hostPort, serviceName),
	)
	if err != nil {
		log.Fatalf("unable to create Zipkin tracer: %v", err)
		return
	}
	opentracing.InitGlobalTracer(tracer)
	return
}

func ServerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		operationName := req.Method()

		//extract metadata to context
		ctx = ContextFromGRPC(ctx, opentracing.GlobalTracer(), operationName)

		//get span from context metadata
		span := opentracing.SpanFromContext(ctx)
		if span == nil {
			//create new root span
			//span = opentracing.StartSpan(operationName)
			return nil
		}

		//span.SetOperationName(operationName)
		defer span.Finish()

		ext.SpanKindRPCServer.Set(span)
		span.SetTag("test tag", "fuck")

		log.Printf("[Trace Wrapper] Before serving request method: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		log.Printf("[Trace Wrapper] After serving request. TraceId: %v\n", opentracing.GlobalTracer())

		return err
	}
}

func ContextFromGRPC(ctx context.Context, tracer opentracing.Tracer, operationName string) context.Context {
	md, _ := metadata.FromContext(ctx)

	var span opentracing.Span
	wireContext, err := tracer.Extract(opentracing.TextMap, metadataReader{&md})
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		log.Printf("metadata error %s\n", err)
	}
	span = tracer.StartSpan(operationName, ext.RPCServerOption(wireContext))
	return opentracing.ContextWithSpan(ctx, span)
}

// A type that conforms to opentracing.TextMapReader and
// opentracing.TextMapWriter.
type metadataReader struct {
	*metadata.MD
}

func (w metadataReader) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range *w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
