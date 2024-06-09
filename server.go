package main

type Config struct {
	Addr  string
	Store Storer
}

type Server struct {
	*Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{Config: cfg}, nil
}
