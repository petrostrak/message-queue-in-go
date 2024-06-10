package main

import (
	"log"
)

// Underlying storage (in memory, on disk, s3)
// Server (http, tcp)

func main() {
	cfg := &Config{
		Addr: ":3000",
		StorerProducerFunc: func() Storer {
			return NewMemoryStore()
		},
	}

	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s.Start()
}
