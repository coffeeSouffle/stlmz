package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unsafe"
)

type Work struct {
	Transactions    int
	Availability    float32
	ElapsedTime     time.Duration
	TotalTransfer   float64
	HTMLTransfer    float32
	TransactionRate float64
	Throughput      float64
	Successful      int
	Failed          int
	StateCode       int

	TransactionTime time.Duration
	ConnectionTime  time.Duration
	ResponseTime    time.Duration

	TotalTransTime time.Duration
	TotalConnTime  time.Duration
	TotalRespTime  time.Duration
}

func (w *Work) Monitor(params Params) {
	fmt.Println(len(request))
	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt)

	timer := new(time.Timer)
	if params.time > 0 {
		timer = time.NewTimer(time.Duration(params.time) * time.Second)
	}

	begin := time.Now()
	defer func() {
		w.ElapsedTime = time.Now().Sub(begin)
		fmt.Println(w.ElapsedTime)
		w.Count(params)
		w.Dump()
		over <- true
	}()

	finish := make(chan bool, params.concurrent)

	for i := 0; i < params.concurrent; i++ {
		go w.Run(params, finish)
	}

	count := 0
	for {
		select {
		case <-finish:
			count += 1
			if count == params.concurrent {
				return
			}
		case <-userInterrupt:
			return
		case <-timer.C:
			return
		}
	}
}

func (w *Work) Dump() {
	fmt.Println(fmt.Sprintf("%s%d%s", "Transactions:                   ", w.Transactions, " hits"))
	fmt.Println(fmt.Sprintf("%s%5.2f%s", "Availability:                   ", w.Availability*100, " %"))
	fmt.Println(fmt.Sprintf("%s%7.2f%s", "Elapsed time:                ", w.ElapsedTime.Seconds(), " secs"))
	fmt.Println(fmt.Sprintf("%s%d%s", "Document length:               ", maxRecvSize, " Bytes"))
	fmt.Println(fmt.Sprintf("%s%8.2f%s", "TotalTransfer:              ", w.TotalTransfer, " MB"))
	fmt.Println(fmt.Sprintf("%s%7.2f%s", "Transaction rate:            ", w.TransactionRate, " trans/sec"))
	fmt.Println(fmt.Sprintf("%s%5.2f%s", "Throughput:                    ", w.Throughput, " MB/sec"))
	fmt.Println(fmt.Sprintf("%s%d%s", "Successful:                     ", w.Successful, " hits"))
	fmt.Println(fmt.Sprintf("%s%d%s", "Failed:                           ", w.Failed, " hits"))
}

func (w *Work) Count(params Params) {
	w.Transactions = params.request
	if w.Transactions != 0 {
		w.Availability = float32(w.Successful) / float32(w.Transactions)
	} else {
		w.Availability = 0
	}

	w.TotalTransfer = float64(maxRecvSize) / 1024

	w.TransactionRate = float64(w.Transactions) / float64(w.ElapsedTime.Seconds())
	w.Throughput = w.TotalTransfer / float64(w.ElapsedTime.Seconds())
}

func (w *Work) Run(params Params, finish chan bool) {
	for {
		select {
		case req := <-request:
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				w.Failed += 1
			} else {
				if res.StatusCode == 200 {
					w.Successful += 1
					maxRecvSize += unsafe.Sizeof(*res)
				} else {
					w.Failed += 1
				}
			}
			if len(request) == 0 {
				finish <- true
				return
			}
		}
	}
}
