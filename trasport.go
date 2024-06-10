package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
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
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)

	if r.Method == "POST" {
		if len(parts) != 2 {
			fmt.Println("invalid action")
			return
		}
		topic := parts[1]
		fmt.Println(topic)
	}

	fmt.Println(parts)
}
