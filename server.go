package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Addr               string
	StorerProducerFunc StorerProducerFunc
}

type Message struct {
	Topic string
	Data  []byte
}

type Server struct {
	*Config
	topics      map[string]Storer
	consumers   []Consumer
	producers   []Producer
	quit        chan struct{}
	produceChan chan Message
}

func NewServer(cfg *Config) (*Server, error) {
	prodChan := make(chan Message)
	return &Server{
		Config:      cfg,
		topics:      make(map[string]Storer),
		producers:   []Producer{NewHTTPProducer(cfg.Addr, prodChan)},
		quit:        make(chan struct{}),
		produceChan: prodChan,
	}, nil
}

func (s *Server) Start() {
	for _, consumer := range s.consumers {
		if err := consumer.Start(); err != nil {
			fmt.Println(err)
		}
	}

	for _, producer := range s.producers {
		go func(p Producer) {
			if err := p.Start(); err != nil {
				fmt.Println(err)
			}
		}(producer)
	}
	s.loop()
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

func (s *Server) loop() {
	for {
		select {
		case <-s.quit:
			return
		case msg := <-s.produceChan:
			fmt.Println("produced msg", msg)
		}
	}
}
