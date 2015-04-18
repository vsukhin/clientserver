package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	PORT_TCP                = 5555
	THREAD_COUNT            = 10
	FILE_NAME               = "data.txt"
	PARAM_NAME_PORT_TCP     = "tcpPort"
	PARAM_NAME_THREAD_COUNT = "threadCount"
	PARAM_NAME_FILE_NAME    = "fileName"
	COMMAND_QUIT            = "quit"
)

var (
	tcpPort     = flag.Int(PARAM_NAME_PORT_TCP, PORT_TCP, "TCP Server port")
	threadCount = flag.Int(PARAM_NAME_THREAD_COUNT, THREAD_COUNT, "Number of parallel threads")
	fileName    = flag.String(PARAM_NAME_FILE_NAME, FILE_NAME, "Name of text data file")
)

func main() {
	Log.Println("Starting client work at ", time.Now())
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	textdata, err := ioutil.ReadFile(*fileName)
	if err != nil {
		Log.Fatalf("Can't read form file %v having error %v", *fileName, err)
	}

	for i := 0; i < *threadCount; i++ {
		tcpClient := NewTcpClient(*tcpPort, string(textdata))
		go tcpClient.Start()
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
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

	Log.Println("Finishing client work at ", time.Now())
}
