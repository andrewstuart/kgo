package main

import (
	"fmt"
	"net/http"

	health "astuart.co/go-healthcheck"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	setupConfig()
	lg.Info("started")

	reg := health.NewRegistry()
	r := mux.NewRouter()
	r.Path("/health").Handler(reg)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Path("/").HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	{{ if .CORS -}}
	http.ListenAndServe(
		":8080", 
		handlers.CORS(
			handlers.AllowedHeaders([]string{"Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(r),
	)
{{ else -}}
	http.ListenAndServe(":8080", r)
{{ end -}}
}
