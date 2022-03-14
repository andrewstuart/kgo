package main

import (
	"fmt"
	"net/http"

	health "astuart.co/go-healthcheck"
	"github.com/andrewstuart/multierrgroup"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gopuff/morecontext"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	meter = global.Meter("{{ .Name | lower }}", metric.WithDescription("time taken for pic download"))
	ctr   = meter.NewFloat64Histogram("hello_world_latency", instrument.WithUnit(unit.Second))
)

func main() {
	ctx := morecontext.ForSignals()
	shutdown := setupOtel(ctx)
	cfg := setupConfig()
	lg.Info("started")

	reg := health.NewRegistry()
	r := mux.NewRouter()
	r.Path("/health").Handler(reg)

	api := r.PathPrefix("/api/v1").Subrouter()

	// Use middleware IRL
	api.Use(mux.MiddlewareFunc(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				ctr.Record(ctx, time.Since(start).Seconds(), attribute.String("path", r.URL.Path))
			}()
			h.ServeHTTP(w, r)
		})
	}))

	api.Path("/").HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		ctx, sp := otel.Tracer("main.go").Start(ctx, "main")
		defer sp.End()

		fmt.Fprintf(w, "hello " + cfg.GetString("name"))
	})

	srv := http.Server{
		Addr: ":8080",
		{{ if .CORS -}}
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(r),
		{{ else -}}
		Handler: r,
		{{ end -}}
	}

	go srv.ListenAndServe()

	<-ctx.Done()

	var meg multierrgroup.Group
	meg.Go(func() error {
		return stop(ctx)
	})
	meg.Go(func() error {
		return srv.Shutdown(ctx)
	})

	err := meg.Wait()
	if err != nil {
		logrus.WithError(err).Fatal("error shutting down gracefully")
	}
}
