package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func setupOtel(ctx context.Context) {
	// Traces
	tr, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("opentelemetry-collector.opentelemetry-collector:4317"), otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("redditpsql"),
			semconv.ServiceVersionKey.String("develop"),
			attribute.String("env", "production"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithBatcher(tr),
	))
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Metrics
	client := otlpmetricgrpc.NewClient(otlpmetricgrpc.WithEndpoint("opentelemetry-collector.opentelemetry-collector:4317"), otlpmetricgrpc.WithInsecure())
	exp, err := otlpmetric.New(ctx, client)
	if err != nil {
		log.Fatalf("Failed to create the collector exporter: %v", err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithHistogramDistribution(),
			exp,
		),
		controller.WithExporter(exp),
		controller.WithCollectPeriod(2*time.Second),
	)
	global.SetMeterProvider(pusher)
}
