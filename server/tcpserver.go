package main

import (
	"io"
	"net"
	"strconv"
	"time"
)

const (
	BUFFER_SIZE = 512
)

type TcpServer struct {
	tcpPort     int
	textChannel chan string
}

func NewTcpServer(tcpport int, textchannel chan string) (tcpserver *TcpServer) {
	tcpserver = new(TcpServer)
	tcpserver.tcpPort = tcpport
	tcpserver.textChannel = textchannel

	return tcpserver
}

func (tcpserver *TcpServer) Start() {
	Log.Println("Starting tcp server at ", time.Now())

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(tcpserver.tcpPort))
	if err != nil {
		Log.Fatalf("Can't listen at port %v having error %v", tcpserver.tcpPort, err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			Log.Fatalf("Error %v during port %v listening", err, tcpserver.tcpPort)
		}
		go tcpserver.Read(conn)
	}
}

func (tcpserver *TcpServer) Read(conn net.Conn) {
	for {
		buf := make([]byte, BUFFER_SIZE)
		count, err := conn.Read(buf)
		if err == io.EOF {
			err = conn.Close()
			if err != nil {
				Log.Fatalf("Can't close connection %v", err)
			}
			return
		}
		if err != nil {
			Log.Fatalf("Error during reading from port %v having error %v", tcpserver.tcpPort, err)
		}
		if count != 0 {
			tcpserver.textChannel <- string(buf)
			Log.Println("Read raw text ", string(buf))
		}
	}
}
