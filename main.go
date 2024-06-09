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
	}

	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
