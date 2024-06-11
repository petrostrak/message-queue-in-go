package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	Addr        string
	ProduceChan chan<- Message
}

func NewHTTPProducer(addr string, produceChan chan Message) *HTTPProducer {
	return &HTTPProducer{
		Addr:        addr,
		ProduceChan: produceChan,
	}
}

func (p *HTTPProducer) Start() error {
	slog.Info("HTTP Trasport started", "port", p.Addr)
	return http.ListenAndServe(p.Addr, p)
}
func (p *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)

	if r.Method == "POST" {
		if len(parts) != 2 {
			fmt.Println("invalid action")
			return
		}
		p.ProduceChan <- Message{
			Topic: parts[1],
			Data:  []byte("some data"),
		}
	}
}
