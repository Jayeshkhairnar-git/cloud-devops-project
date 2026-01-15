package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	
)

func InitTracer(serviceName string) (func(context.Context) error, error) {
	ctx := context.Background()

	// OTLP exporter (Cloud Run → Google Cloud Trace)
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("cloudtrace.googleapis.com:443"),
	)
	if err != nil {
		return nil, err
	}

	// Resource defines service.name (THIS is what shows in Trace Explorer)
	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.DeploymentEnvironment("cloudrun"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(
			exporter,
			sdktrace.WithBatchTimeout(5*time.Second),
		),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}


this is go code suggest me changes so that my teams name come in telementary service
