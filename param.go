package main

import (
	"flag"
	"fmt"
	// cli "github.com/codegangsta/cli"
	"os"
)

type Params struct {
	concurrent int // number of concurrent
	request    int // number of request
	time       int // testing time, 30 mean 30seconds
	reps       int // number of times to run the testing

	contentType string
	httpVersion string
	urlFile     string
	postFile    string
}

func (p *Params) ParseCommandLine() {
	p.parseHelpCmd()

	fs := flag.NewFlagSet("stlmz", flag.ExitOnError)
	fs.IntVar(&p.concurrent, "c", 1, "number of concurrent")
	fs.IntVar(&p.concurrent, "concurrent", 1, "number of concurrent")
	fs.IntVar(&p.request, "r", 1, "number of request")
	fs.IntVar(&p.request, "request", 1, "number of request")
	fs.IntVar(&p.time, "t", 0, "execute testing time (seconds)")
	fs.IntVar(&p.time, "time", 0, "execute testing time (seconds)")
	fs.IntVar(&p.reps, "R", 1, "execute times")
	fs.IntVar(&p.reps, "repetitions", 1, "execute times")
	fs.StringVar(&p.contentType, "T", "", "set Content-Type in request")
	fs.StringVar(&p.contentType, "content-type", "", "set Content-Type in request")
	fs.StringVar(&p.httpVersion, "v", "", "set http version")
	fs.StringVar(&p.httpVersion, "http-version", "", "set http version")
	fs.Parse(os.Args[1:])

	p.checkParams()

	fmt.Println(fs.Args())
}

func (p *Params) checkParams() {
	if p.concurrent == 0 {
		howtoUse()
	}

	if p.request == 0 {
		howtoUse()
	}

	if p.reps == 0 {
		howtoUse()
	}
}

func (p *Params) parseHelpCmd() {

	if len(os.Args) == 1 {
		howtoUse()
	}

	for _, v := range os.Args {
		switch v {
		case "-h":
			howtoUse()
		case "-V":
			GetVersion()
		}
	}
}

func GetVersion() {
	fmt.Println("stlmz Version: ", Version)
	os.Exit(2)
}

func howtoUse() {
	fmt.Println(HelpInfo)
	os.Exit(2)
}

func InitParams(params *Params) {
	params.ParseCommandLine()
}
