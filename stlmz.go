package main

import (
	"fmt"
	"net/http"
	// "os"
	"runtime"
)

var over chan bool
var work Work
var concurrents chan bool
var request chan *http.Request
var requestflag chan bool

var maxRecvSize uintptr

func main() {
	params := Params{}
	InitParams(&params)
	prepareData(params)

	fmt.Println(params)

	<-requestflag
	go work.Monitor(params)

	<-over
}

func prepareData(params Params) {
	over = make(chan bool)
	requestflag = make(chan bool)
	concurrents = make(chan bool, params.concurrent)
	request = make(chan *http.Request, params.request)
	work = Work{}
	runtime.GOMAXPROCS(1) //runtime.NumCPU()
}
