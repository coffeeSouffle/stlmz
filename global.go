package main

var Version string = "0.0.0.1"

const (
	HelpInfo = `
Usage:
    stlmz [options] [http method] http[s]://hostname[:post][/uri]

Options:
    -c, --concurrent     concurrent users. (default: 1)
    -r, --requests       execute number of all requests. (default: 1)
    -R, --repetitions    number of times to run the testing. (default: 1)
    -t, --time           testing execute time. (unit: seconds)
    -f, --file           file have request what need params.
    -v, --http-version   set http-version. (default: HTTP/1.1)

    -V, --Version        stlmz version
    -h, --help           stlmz help

Example:
    stlmz -c 10 -n 100 GET http://www.google.com/

Copyright (C) 2015 by LMZ.
This is free software; open source on github.com/coffeeSouffle/stlmz.
    `
)
