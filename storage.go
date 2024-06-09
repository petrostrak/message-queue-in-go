package main

import "fmt"

type Storer interface {
	Push([]byte) error
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) error {
	s.data = append(s.data, b)
	return nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}

	return s.data[offset], nil
}
