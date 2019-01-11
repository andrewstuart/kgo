package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	setupConfig()
	lg.Info("started")

	r := mux.NewRouter()

	r.Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.ListenAndServe(":8080", r)
}
