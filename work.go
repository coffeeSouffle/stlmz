package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var maxRecvSize int

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
		w.Dump()
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
	fmt.Println(fmt.Sprintf("%s%8.3f%s", "TransactionTime:            ", w.TransactionTime.Seconds()*1000, " ms(mean)"))
	fmt.Println(fmt.Sprintf("%s%8.3f%s", "ConnectionTime:             ", w.ConnectionTime.Seconds()*1000, " ms(mean)"))
	fmt.Println(fmt.Sprintf("%s%8.3f%s", "ProcessTime:                ", w.ResponseTime.Seconds()*1000, " ms(mean)"))
	fmt.Println(fmt.Sprintf("%s%d%s", "StateCode:                    ", w.StateCode, "(code 200)"))
}
