package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Params struct {
	concurrent int // number of concurrent
	request    int // number of request
	time       int // testing time, 30 mean 30seconds
	reps       int // number of times to run the testing

	user        string
	method      string
	httpVersion string
	url         string
	param       url.Values
	file        string
	header      map[string]string
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
	fs.StringVar(&p.file, "f", "", "file have request what need params")
	fs.StringVar(&p.file, "file", "", "file have request what need params")
	fs.StringVar(&p.httpVersion, "v", "", "set http version")
	fs.StringVar(&p.httpVersion, "http-version", "", "set http version")
	fs.Parse(os.Args[1:])

	p.checkParams()

	args := fs.Args()
	if len(args) < 2 {
		howtoUse()
	}

	p.method = args[0]
	p.url = args[1]

	if p.file != "" {
		var dat map[string]map[string]string
		fptr, err := ioutil.ReadFile(p.file)
		if err != nil {
			howtoUseFile(err)
		}

		if err := json.Unmarshal(fptr, &dat); err != nil {
			howtoUseFile(err)
		}

		fmt.Println(dat)
		p.parseJson(dat)
	}

	go p.SetRequest()
}

func (p *Params) parseJson(data map[string]map[string]string) {
	pm := url.Values{}
	for k, v := range data {
		// fmt.Println(v)
		switch k {
		case "header":
			p.header = v
		case "param":
			for key, val := range v {
				pm.Add(key, val)
			}
			p.param = pm
		}
	}
}

func (p *Params) checkParams() {
	if p.concurrent == 0 {
		howtoUse()
	}

	if p.request < p.concurrent {
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

func (p *Params) SetRequest() {
	requestflag <- true
	for i := 0; i < p.request; i++ {
		req, err := http.NewRequest(p.method, p.url, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if len(p.header) > 0 {
			for k, v := range p.header {
				req.Header.Add(k, v)
			}
		}

		if p.method == "POST" {
			req.Header.Add("Content-Length", strconv.Itoa(len(p.param.Encode())))
		}

		if req != nil {
			request <- req
		}
	}
	fmt.Println(len(request))
}

func GetVersion() {
	fmt.Println("stlmz Version: ", Version)
	os.Exit(2)
}

func howtoUse() {
	fmt.Println(HelpInfo)
	os.Exit(2)
}

// TODO 作提示
func howtoUseFile(i error) {
	log.Fatal("file error: ", i)
	os.Exit(2)
}

func InitParams(params *Params) {
	params.ParseCommandLine()
}
