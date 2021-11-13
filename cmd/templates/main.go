package main

import (
	"fmt"
	"net/http"

	health "astuart.co/go-healthcheck"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	cfg := setupConfig()
	lg.Info("started")

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		logrus.WithError(err).Fatal("couldn't load tracer config")
	}
	tracer, closer, err := cfg.InitGlobalTracer("{{ .Name }}")

	reg := health.NewRegistry()
	r := mux.NewRouter()
	r.Path("/health").Handler(reg)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Path("/").HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	{{ if .CORS -}}
	lg.Fatal(http.ListenAndServe(
		":8080", 
		handlers.CORS(
			handlers.AllowedHeaders([]string{"Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(r),
	))
{{ else -}}
	lg.Fatal(http.ListenAndServe(":8080", r))
{{ end -}}
}
