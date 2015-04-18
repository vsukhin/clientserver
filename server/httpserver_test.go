package main

import (
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	var httpPort = 1
	var statistics = new(Storage)

	httpServer := NewHttpServer(httpPort, statistics)
	if httpServer == nil {
		t.Error("New http server should create a new object")
	} else {
		if httpServer.httpPort != httpPort {
			t.Error("New http server should use a right http port")
		}
		if httpServer.statistics != statistics {
			t.Error("New http server should use a right statistics interface")
		}
	}
}
