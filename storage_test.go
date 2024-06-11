package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestStorage(t *testing.T) {
	payload := []byte("a random payload")
	resp, err := http.Post("http://localhost:3000/publish/topic_1", "application/octet-stream", bytes.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200 but got %d\n", resp.StatusCode)
	}
}
