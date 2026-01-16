package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
    //semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	
)

func InitTracer(serviceName string) (func(context.Context) error, error) {
	ctx := context.Background()


	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("cloudtrace.googleapis.com:443"),
	)
	if err != nil {
		return nil, err
	}


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

