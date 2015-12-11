# stlmz

stlmz is a stress testing example with golang.

## Install

```bash
    go get github.com/coffeeSouffle/stlmz
    go install github.com/coffeeSouffle/stlmz
    stlmz -r 100 -c 10 GET http://www.google.com
```

## Help Command

```
Usage:
    stlmz [options] [http method] http[s]://hostname[:post][/uri]

Options:
    -c, --concurrent     concurrent users. (default: 1)
    -r, --requests       execute number of all requests. (default: 1)
    -t, --time           testing execute time. (unit: seconds)
    -f, --file           file have request what need params.
    -v, --http-version   set http-version. (default: HTTP/1.1)

    -V, --Version        stlmz version
    -h, --help           stlmz help

Example:
    stlmz -c 10 -r 100 GET http://www.google.com/

Copyright (C) 2015 by LMZ.
This is free software; open source on github.com/coffeeSouffle/stlmz.
```