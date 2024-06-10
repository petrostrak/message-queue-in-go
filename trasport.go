package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Consumer interface {
	Start() error
}

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	Addr string
}

func NewHTTPProducer(addr string) *HTTPProducer {
	return &HTTPProducer{
		Addr: addr,
	}
}

func (p *HTTPProducer) Start() error {
	slog.Info("HTTP Trasport started", "port", p.Addr)
	return http.ListenAndServe(p.Addr, p)
}
func (p *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
