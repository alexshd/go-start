// Package web contains http related functionality
package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *ShdServer) handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("Hello"))
	}
}

func (s *ShdServer) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		routes := []string{}
		for _, r := range s.route.Routes() {
			routes = append(routes, r.Pattern)
		}

		w.Header().Set("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(routes)
	}
}

func (s *ShdServer) routes() {
	s.route.Get("/", s.handleRoot())
	s.route.Get("/status", s.handleStatus())
}

type ShdServer struct {
	route *chi.Mux
	Port  string
}

func NewServer(host string) *ShdServer {
	s := new(ShdServer)

	s.route = chi.NewRouter()
	s.route.Use(middleware.Logger)
	s.routes()
	s.Port = host

	return s
}

func (s *ShdServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.route.ServeHTTP(w, r)
}
