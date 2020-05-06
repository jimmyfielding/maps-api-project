package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heptiolabs/healthcheck"
)

func (s *server) routes() http.Handler {
	r := mux.NewRouter()
	health := healthcheck.NewHandler()
	r.HandleFunc("/healthz", health.ReadyEndpoint).
		Methods("GET").
		Schemes("http")

	r.HandleFunc("/titles", s.getTitles).
		Methods("GET").
		Schemes("http")

	return r
}
