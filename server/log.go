package main

import (
	"log"
)

type Logger interface {
	Println(args ...interface{})
	Fatalf(query string, args ...interface{})
}

type DefaultLogger struct {
}

var (
	Log Logger = new(DefaultLogger)
)

func (logger *DefaultLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (logger *DefaultLogger) Fatalf(query string, args ...interface{}) {
	log.Fatalf(query, args...)
}
