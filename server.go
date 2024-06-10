package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Addr               string
	StorerProducerFunc StorerProducerFunc
}

type Server struct {
	*Config
	topics map[string]Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
		topics: make(map[string]Storer),
	}, nil
}

func (s *Server) Start() {
	http.ListenAndServe(s.Addr, s)
}

func (s *Server) createTopic(name string) bool {
	if _, found := s.topics[name]; !found {
		s.topics[name] = s.StorerProducerFunc()
		return true
	}

	return false
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
