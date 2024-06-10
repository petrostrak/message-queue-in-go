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
	topics    map[string]Storer
	consumers []Consumer
	quit      chan struct{}
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
		topics: make(map[string]Storer),
		quit:   make(chan struct{}),
	}, nil
}

func (s *Server) Start() {
	for _, consumer := range s.consumers {
		if err := consumer.Start(); err != nil {
			fmt.Println(err)
		}
	}
	<-s.quit
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
