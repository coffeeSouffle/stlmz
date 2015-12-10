package main

import (
	"fmt"
	"runtime"
)

var over chan bool
var work Work
var concurrents chan bool

func main() {
	params := Params{}
	InitParams(&params)
	prepareData(params)

	fmt.Println(params)

	go work.Monitor(params)
	<-over
}

func prepareData(params Params) {
	over = make(chan bool)
	concurrents = make(chan bool, params.concurrent)
	work = Work{}
	runtime.GOMAXPROCS(1) //runtime.NumCPU()
}
