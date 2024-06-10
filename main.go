package main

import (
	"log"
)

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
