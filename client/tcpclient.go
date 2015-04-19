package main

import (
	"net"
	"strconv"
	"time"
)

type TcpClient struct {
	tcpPort  int
	textData string
}

func NewTcpClient(tcpport int, textdata string) (tcpclient *TcpClient) {
	tcpclient = new(TcpClient)
	tcpclient.tcpPort = tcpport
	tcpclient.textData = textdata

	return tcpclient
}

func (tcpclient *TcpClient) Start() {
	Log.Println("Starting tcp client at ", time.Now())
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(tcpclient.tcpPort))
	if err != nil {
		Log.Fatalf("Can't connect to port %v having error %v", tcpclient.tcpPort, err)
	}
	defer conn.Close()

	for {
		_, err = conn.Write([]byte(tcpclient.textData))
		if err == nil {
			Log.Println("Succesfully sending text data ", tcpclient.textData)
		} else {
			Log.Fatalf("Can't send text data %v having error %v", tcpclient.textData, err)
		}
	}
}
