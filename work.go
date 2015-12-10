package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Work struct {
	Transactions    int
	Availability    float32
	ElapsedTime     time.Duration
	TotalTransfer   float64
	HTMLTransfer    float32
	TransactionRate float32
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
		over <- true
	}()

	for {
		select {
		case <-userInterrupt:
			return
		case <-timer.C:
			return
		}
	}
}
