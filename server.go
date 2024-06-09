package main

type Config struct {
	Addr string
}

type Server struct {
	*Config
	Store Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{Config: cfg}, nil
}
