package main

import (
	"testing"
)

func TestNewTcpServer(t *testing.T) {
	var tcpport int = 1
	var channel = make(chan string, 1)

	tcpserver := NewTcpServer(tcpport, channel)
	if tcpserver == nil {
		t.Error("New tcp server should create a new object")
	} else {
		if tcpserver.tcpPort != tcpport {
			t.Error("New tcp server should use a right tcp port value")
		}
		if cap(tcpserver.textChannel) != 1 {
			t.Error("New tcp server should use a right text channel")
		}
	}
}
