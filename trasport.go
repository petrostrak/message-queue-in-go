package main

type Consumer interface {
	Start() error
}

type Producer interface{}
