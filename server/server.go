package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	PORT_TCP              = 5555
	PORT_HTTP             = 8080
	PARAM_NAME_STORE_TEXT = "storeText"
	PARAM_NAME_PORT_TCP   = "tcpPort"
	PARAM_NAME_PORT_HTTP  = "httpPort"
	COMMAND_QUIT          = "quit"
	CAPACITY_CHANNEL      = 100
)

var (
	storeText = flag.Bool(PARAM_NAME_STORE_TEXT, false, "Store original raw text")
	tcpPort   = flag.Int(PARAM_NAME_PORT_TCP, PORT_TCP, "TCP Server port")
	httpPort  = flag.Int(PARAM_NAME_PORT_HTTP, PORT_HTTP, "HTTP server port")
)

func main() {
	Log.Println("Starting server work at ", time.Now())
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	DataStorage := NewStorage(*storeText)
	textChannel := make(chan string, CAPACITY_CHANNEL)
	tcpServer := NewTcpServer(*tcpPort, textChannel)
	httpServer := NewHttpServer(*httpPort, DataStorage)

	go DataStorage.ProcessText(textChannel)
	go tcpServer.Start()
	go httpServer.Start()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			Log.Fatalf("Can't read user input %v", err)
		}
		if strings.ToLower(strings.TrimSpace(command)) == COMMAND_QUIT {
			break
		} else {
			fmt.Println("Uknown command. List of supported commands: ", COMMAND_QUIT)
		}
	}

	Log.Println("Finishing server work at ", time.Now())
}
