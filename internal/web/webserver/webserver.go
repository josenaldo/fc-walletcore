package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebSerber struct {
	Router         chi.Router
	Handlers       map[string]http.HandlerFunc
	WerbServerPort string
}

func NewWebServer(
	webServerPort string) *WebSerber {
	return &WebSerber{
		Router:         chi.NewRouter(),
		Handlers:       make(map[string]http.HandlerFunc),
		WerbServerPort: webServerPort,
	}
}

func (s *WebSerber) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebSerber) Start() {
	s.Router.Use(middleware.Logger)

	for path, handler := range s.Handlers {
		s.Router.Post(path, handler)
	}

	http.ListenAndServe(s.WerbServerPort, s.Router)
}
