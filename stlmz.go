package main

import (
	"fmt"
	"net/http"
	"runtime"
)

var over chan bool
var work Work
var concurrents chan bool

//http://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest

func main() {
	params := Params{}
	InitParams(&params)
	prepareData(params)

	fmt.Println(params)

	go work.Monitor(params)

	client := &http.Client{}
	req, err := http.NewRequest(params.method, params.url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		if res.StatusCode == 200 {
			work.Successful += 1
		} else {
			work.Failed += 1
		}
	}

	<-over
}

func prepareData(params Params) {
	over = make(chan bool)
	concurrents = make(chan bool, params.concurrent)
	work = Work{}
	runtime.GOMAXPROCS(1) //runtime.NumCPU()
}
