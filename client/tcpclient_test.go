package main

import (
	"testing"
)

func TestNewTcpClient(t *testing.T) {
	var tcpport int = 1
	var textdata = "text"

	tcpclient := NewTcpClient(tcpport, textdata)
	if tcpclient == nil {
		t.Error("New tcp client should create a new object")
	} else {
		if tcpclient.tcpPort != tcpport {
			t.Error("New tcp client should use a right tcp port value")
		}
		if tcpclient.textData != textdata {
			t.Error("New tcp client should use a right text data")
		}
	}
}
