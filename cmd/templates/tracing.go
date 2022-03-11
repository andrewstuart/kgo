package main

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func init() {
	tr, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("{{ .Otel.CollectorEndpoint }}"), otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("{{ .Name }}"),
			semconv.ServiceVersionKey.String("develop"),
			attribute.String("env", "prod"),
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
}
