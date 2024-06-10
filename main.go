package main

import (
	"fmt"
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

	s.Store.Push([]byte("foobar"))
	data, err := s.Store.Fetch(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
