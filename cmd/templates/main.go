package main

import (
	"fmt"
	"net/http"

	health "astuart.co/go-healthcheck"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gopuff/morecontext"
)

var (
	meter = global.Meter("{{ .Name | tolower }}", metric.WithDescription("time taken for pic download"))
	ctr   = meter.NewFloat64Histogram("hello_world_latency", instrument.WithUnit(unit.Second))
)

func main() {
	ctx := morecontext.ForSignals()
	setupOtel(ctx)
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

}
