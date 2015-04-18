package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
	httpPort   int
	statistics Statistics
}

func NewHttpServer(httpport int, statistics Statistics) (httpServer *HttpServer) {
	httpServer = new(HttpServer)
	httpServer.httpPort = httpport
	httpServer.statistics = statistics

	return httpServer
}

func (httpserver *HttpServer) Start() {
	http.HandleFunc("/", httpserver.HomePage)
	http.Handle("/stats", httpserver)

	Log.Println("Starting http server at ", time.Now())

	err := http.ListenAndServe("localhost:"+strconv.Itoa(httpserver.httpPort), nil)
	if err != nil {
		Log.Fatalf("Can't listen http port %v having error %v", httpserver.httpPort, err)
	}
}

func (httpserver *HttpServer) HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Please use /stats page for statistics information")
}

func (httpserver *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	results, err := json.Marshal(httpserver.statistics.GetResults())
	if err != nil {
		Log.Fatalf("Can't marshal results %v to json %v format", httpserver.statistics.GetResults(), err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(results))
}
